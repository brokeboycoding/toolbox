package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var aStr string
var bStr string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Cong 2 so nguyen(int va float)",

	Run: func(cmd *cobra.Command, args []string) {
		//chuyen doi string sang float64
		a, err1 := strconv.ParseFloat(aStr, 64)
		b, err2 := strconv.ParseFloat(bStr, 64)

		if err1 != nil || err2 != nil {
			fmt.Println("❌ Vui lòng nhập số hợp lệ cho --a và --b")
			return
		}
		sum := a + b
		fmt.Printf("✅ Tổng của %g và %g là: %g\n", a, b, sum)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&aStr, "a", "x", "", "So thu nhat")
	addCmd.Flags().StringVarP(&bStr, "b", "y", "", "So thu hai")

	//Bat buoc truyen
	addCmd.MarkFlagRequired("a")
	addCmd.MarkFlagRequired("b")
}
