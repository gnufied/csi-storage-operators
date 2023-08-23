package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gnufied/csi-storage-operators/pkg/operator"
	"github.com/openshift/library-go/pkg/controller/controllercmd"
	"github.com/spf13/cobra"
	"k8s.io/component-base/cli"
	"k8s.io/component-base/version"
)

var guestKubeconfig *string

func main() {
	fmt.Printf("Hello World")
	command := NewOperatorCommand()
	code := cli.Run(command)
	os.Exit(code)
}

func NewOperatorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "csi-storage-operator",
		Short: "OpenShift CSI Driver Operator",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}

	ctrlCmd := controllercmd.NewControllerCommandConfig(
		"csi-storage-operator",
		version.Get(),
		runOperatorWithGuestKubeconfig,
	).NewCommand()

	guestKubeconfig = ctrlCmd.Flags().String("guest-kubeconfig", "", "Path to the guest kubeconfig file. This flag enables hypershift integration.")

	ctrlCmd.Use = "start"
	ctrlCmd.Short = "Start the CSI Driver Operator"

	cmd.AddCommand(ctrlCmd)

	return cmd
}

func runOperatorWithGuestKubeconfig(ctx context.Context, controllerConfig *controllercmd.ControllerContext) error {
	return operator.RunOperator(ctx, controllerConfig, *guestKubeconfig)
}
