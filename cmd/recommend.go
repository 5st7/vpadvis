package cmd

import (
	"github.com/5st7/vpadvis/application"
	"github.com/5st7/vpadvis/infrastructure"
	"github.com/5st7/vpadvis/interfaces"
	"github.com/5st7/vpadvis/repository"
	"github.com/spf13/cobra"
)

var format, vpaName, namespace string

var recommendCmd = &cobra.Command{
	Use:   "recommend",
	Short: "Display resource recommendations for all VPA in the specified namespace",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Kubernetesクライアントの初期化
		k8sClient, err := infrastructure.NewK8sClient()
		if err != nil {
			return err
		}

		// 各ワークロード用のリポジトリを初期化
		deploymentRepo := repository.NewDeploymentRepository(k8sClient)
		statefulSetRepo := repository.NewStatefulSetRepository(k8sClient)
		daemonSetRepo := repository.NewDaemonSetRepository(k8sClient)

		// VPAリポジトリの作成
		vpaRepo := repository.NewVPARepository(k8sClient, deploymentRepo, statefulSetRepo, daemonSetRepo)
		vpaService := application.NewVPAService(vpaRepo)

		// ネームスペース内のすべてのVPAリソースの推奨設定の取得
		workloadResources, err := vpaService.GetAllVPARecommendations(namespace)
		if err != nil {
			return err
		}

		// フォーマッターの選択と出力
		formatter := interfaces.NewFormatter(format)
		formatter.PrintAllRecommendations(workloadResources)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(recommendCmd)
	recommendCmd.Flags().StringVarP(&format, "format", "f", "markdown", "Output format (markdown, plaintext)")
	recommendCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "Namespace")
}
