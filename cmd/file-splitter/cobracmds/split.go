package cobracmds

import (
	"os"

	filesplitter "github.com/ross96D/file-splitter"
	"github.com/spf13/cobra"
)

var sizeSplit uint64

var splitCommand cobra.Command = cobra.Command{
	Use:   "split",
	Short: "Split the select file into pieces",
	Long: `With this command you can split a file into pieces of a N size

fsplit split -o <output-file-path> -s <size-of-the-file> <path-of-the-file-to-be-splitted>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open(args[0])
		if err != nil {
			panic(err)
		}
		println("SSSSSSSSs")
		defer f.Close()
		err = filesplitter.SplitFile(f, output, int64(sizeSplit))
		if err != nil {
			println("ERR")
			panic(err)
		}
		println("SSSSSSSSs")
	},
}

func init() {
	splitCommand.Flags().Uint64VarP(&sizeSplit, "size", "s", 20*(1<<20), "size of the splitted parts in bytes. Default 20 megas")
}
