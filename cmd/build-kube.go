package cmd

import (
	"github.com/SUSE/fissile/kube"
	"github.com/SUSE/fissile/model"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	flagBuildKubeOutputDir       string
	flagBuildKubeDefaultEnvFiles []string
	flagBuildKubeUseMemoryLimits bool
	flagBuildKubeTagExtra        string
)

// buildKubeCmd represents the kube command
var buildKubeCmd = &cobra.Command{
	Use:   "kube",
	Short: "Creates Kubernetes configuration files.",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		flagBuildKubeOutputDir = buildKubeViper.GetString("output-dir")
		flagBuildKubeDefaultEnvFiles = splitNonEmpty(buildKubeViper.GetString("defaults-file"), ",")
		flagBuildKubeUseMemoryLimits = buildKubeViper.GetBool("use-memory-limits")
		flagBuildKubeTagExtra = buildKubeViper.GetString("tag-extra")

		err := fissile.LoadReleases(
			flagRelease,
			flagReleaseName,
			flagReleaseVersion,
			flagCacheDir,
		)
		if err != nil {
			return err
		}

		opinions, err := model.NewOpinions(
			flagLightOpinions,
			flagDarkOpinions,
		)
		if err != nil {
			return err
		}

		settings := &kube.ExportSettings{
			OutputDir:       flagBuildKubeOutputDir,
			Registry:        flagDockerRegistry,
			Username:        flagDockerUsername,
			Password:        flagDockerPassword,
			Organization:    flagDockerOrganization,
			Repository:      flagRepository,
			UseMemoryLimits: flagBuildKubeUseMemoryLimits,
			FissileVersion:  fissile.Version,
			Opinions:        opinions,
			CreateHelmChart: false,
			TagExtra:        flagBuildKubeTagExtra,
		}

		return fissile.GenerateKube(flagRoleManifest, flagBuildKubeDefaultEnvFiles, settings)
	},
}
var buildKubeViper = viper.New()

func init() {
	initViper(buildKubeViper)

	buildCmd.AddCommand(buildKubeCmd)

	buildKubeCmd.PersistentFlags().StringP(
		"output-dir",
		"",
		".",
		"Kubernetes configuration files will be written to this directory",
	)

	buildKubeCmd.PersistentFlags().StringP(
		"defaults-file",
		"D",
		"",
		"Env files that contain defaults for the parameters generated by kube",
	)

	buildKubeCmd.PersistentFlags().BoolP(
		"use-memory-limits",
		"",
		true,
		"Include memory limits when generating kube configurations",
	)

	buildKubeCmd.PersistentFlags().StringP(
		"tag-extra",
		"",
		"",
		"Additional information to use in computing the image tags",
	)

	buildKubeViper.BindPFlags(buildKubeCmd.PersistentFlags())
}
