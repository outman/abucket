package cmd

/*
Copyright © 2020 pochonlee@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

import (
	"github.com/outman/abucket/internal/model"
	"github.com/outman/abucket/internal/pkg"
	"github.com/spf13/cobra"
)

// automigrateCmd represents the automigrate command
var automigrateCmd = &cobra.Command{
	Use:   "automigrate",
	Short: "Auto migrate database tables.",
	Long: `AutoMigrate will ONLY create tables, 
			missing columns and missing indexes, 
			and WON’T change existing column’s type 
			or delete unused columns to protect your data.`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.NewMySQL().DB.AutoMigrate(&model.Experiment{}, &model.Layer{})
	},
}

func init() {
	rootCmd.AddCommand(automigrateCmd)
}
