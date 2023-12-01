package cobracmds

import (
	"os"

	filesplitter "github.com/ross96D/file-splitter"
	"github.com/spf13/cobra"
)

var joinCommand cobra.Command = cobra.Command{
	Use:   "join",
	Short: "Split the select file into pieces",
	Long: `With this command you can join a partitioned file. The partitioned selected must finish with .part0

fsplit join -o <output-file-path> -s <size-of-the-file> <path-of-the-file-to-be-join.part0>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Create(output)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		filesplitter.JoinFile(args[0], f)
	},
}
