# gcp-iam-lookup

Search Google Cloud IAM roles that contain specific permissions.

## Install

`go install github.com/joemiller/gcp-iam-lookup@latest`

## Usage

Requires Application Default Credentials. If running localy with `gcloud` you can run `gcloud auth application-default login` to setup local ADC.

```

List all roles containing permission `compute.instances.get`:

```console
$ gcp-iam-lookup compute.instances.get

==> Searching for roles containing permissions: [compute.instances.get] ...
roles/owner (5490)
roles/editor (5063)
roles/viewer (2474)
roles/composer.serviceAgent (1285)
roles/container.serviceAgent (1198)
roles/dataflow.serviceAgent (877)
roles/cloudtpu.serviceAgent (672)
roles/notebooks.legacyAdmin (653)
roles/compute.admin (599)
roles/securitycenter.controlServiceAgent (590)
roles/securitycenter.serviceAgent (589)
roles/compute.networkAdmin (486)
...
```

The number of permissions included in each role is listed in parantheeses.

Roles are sorted from most to least permissions.

List all roles containing two or more permissions. Only roles containing all permissions are returned:

```console
$ gcp-iam-lookup \
  compute.instances.get \
  compute.instances.delete \
  compute.instances.create

==> Searching for roles containing permissions: [compute.instances.get compute.instances.delete compute.instances.create] ...
roles/owner (5490)
roles/editor (5063)
roles/composer.serviceAgent (1285)
roles/container.serviceAgent (1198)
roles/dataflow.serviceAgent (877)
roles/cloudtpu.serviceAgent (672)
roles/notebooks.legacyAdmin (653)
roles/compute.admin (599)
roles/notebooks.serviceAgent (400)
roles/genomics.serviceAgent (288)
roles/lifesciences.serviceAgent (288)
roles/compute.instanceAdmin.v1 (285)
roles/dataproc.serviceAgent (253)
roles/compute.instanceAdmin (166)
roles/appengineflex.serviceAgent (157)
roles/vpcaccess.serviceAgent (94)
roles/cloudmigration.inframanager (88)
```