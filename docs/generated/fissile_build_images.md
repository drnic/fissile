## fissile build images

Builds Docker images from your BOSH releases.

### Synopsis


This command goes through all the instance group definitions in the role manifest creating a
Dockerfile for each of them and building it.

Each instance group gets a directory `<work-dir>/dockerfiles`. In each directory one can find
a Dockerfile and a directory structure that gets ADDed to the docker image. The
directory structure contains jobs, packages and all other necessary scripts and
templates.

The images will have a 'instance_group' label useful for filtering.
The entrypoint for each image is `/opt/fissile/run.sh`.

The images will be tagged: `<repository>-<instance_group_name>:<SIGNATURE>`.
The SIGNATURE is based on the hashes of all jobs and packages that are included in
the image.

The `--patch-properties-release` flag is used to distinguish the patchProperties release/job spec
from other specs.  At most one is allowed.
	

```
fissile build images [flags]
```

### Options

```
      --add-label strings                 Additional label which will be set for the base layer image. Format: label=value
  -F, --force                             If specified, image creation will proceed even when images already exist.
  -h, --help                              help for images
  -N, --no-build                          If specified, the Dockerfile and assets will be created, but the image won't be built.
  -O, --output-directory string           Output the result as tar files in the given directory rather than building with docker
  -P, --patch-properties-release string   Used to designate a "patch-properties" psuedo-job in a particular release.  Format: RELEASE/JOB.
      --roles string                      Build only images with the given instance group name; comma separated.
  -s, --stemcell string                   The source stemcell
      --stemcell-id string                Docker image ID for the stemcell (intended for CI)
      --tag-extra string                  Additional information to use in computing the image tags
```

### Options inherited from parent commands

```
  -c, --cache-dir string             Local BOSH cache directory. (default "~/.bosh/cache")
      --config string                config file (default is $HOME/.fissile.yaml)
  -d, --dark-opinions string         Path to a BOSH deployment manifest file that contains properties that should not have opinionated defaults.
      --docker-organization string   Docker organization used when referencing image names
      --docker-password string       Password for authenticated docker registry
      --docker-registry string       Docker registry used when referencing image names
      --docker-username string       Username for authenticated docker registry
  -l, --light-opinions string        Path to a BOSH deployment manifest file that contains properties to be used as defaults.
  -M, --metrics string               Path to a CSV file to store timing metrics into.
  -o, --output string                Choose output format, one of human, json, or yaml (currently only for 'show properties') (default "human")
      --output-graph string          Output a graphviz graph to the given file name
  -r, --release string               Path to final or dev BOSH release(s).
  -n, --release-name string          Name of a dev BOSH release; if empty, default configured dev release name will be used; Final release always use the name in release.MF
  -v, --release-version string       Version of a dev BOSH release; if empty, the latest dev release will be used; Final release always use the version in release.MF
  -p, --repository string            Repository name prefix used to create image names. (default "fissile")
  -m, --role-manifest string         Path to a yaml file that details which jobs are used for each instance group.
  -V, --verbose                      Enable verbose output.
  -w, --work-dir string              Path to the location of the work directory. (default "/var/fissile")
  -W, --workers int                  Number of workers to use; zero means determine based on CPU count.
```

### SEE ALSO

* [fissile build](fissile_build.md)	 - Has subcommands to build all images and necessary artifacts.

###### Auto generated by spf13/cobra on 29-Nov-2018
