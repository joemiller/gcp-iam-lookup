package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"golang.org/x/exp/slices"
	"google.golang.org/api/iam/v1"
)

type role struct {
	name  string
	perms int
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] iam.permission1 [iam.permission2 ...]\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	perms := os.Args[1:]

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := iam.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stderr, "==> Searching for roles containing permissions: %s ...\n", perms)
	matchingRoles := []role{}

	// View("FULL") instructs the API to include the IncludedPermissions field for each role
	err = client.Roles.List().View("FULL").PageSize(1000).Pages(ctx,
		func(page *iam.ListRolesResponse) error {
			for _, r := range page.Roles {
				matches := 0
				for _, perm := range perms {
					if slices.Contains(r.IncludedPermissions, perm) {
						matches += 1
					}
				}
				if matches == len(perms) {
					matchingRoles = append(matchingRoles, role{r.Name, len(r.IncludedPermissions)})
				}
			}
			return nil
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(matchingRoles, func(a, b int) bool {
		return matchingRoles[a].perms > matchingRoles[b].perms
	})
	for _, role := range matchingRoles {
		fmt.Printf("%s (%d)\n", role.name, role.perms)
	}
}
