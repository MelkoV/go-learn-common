package cmd

import (
	"errors"
	"fmt"
	"github.com/MelkoV/go-learn-common/dictionary/source/admin"
	"github.com/MelkoV/go-learn-common/dictionary/source/user"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

// dictionaryCmd represents the dictionary command
var dictionaryCmd = &cobra.Command{
	Use:   "dictionary",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		lang := viper.GetStringSlice("dictionary.languages")
		for _, v := range lang {
			processLang(v)
		}
	},
}

func processLang(l string) {
	data := map[string]string{
		"user":  processVars(l, "user", user.Values),
		"admin": processVars(l, "admin", admin.Values),
	}
	vars := make([]string, len(data))
	i := 0
	for k, v := range data {
		vars[i] = fmt.Sprintf("\t\"%s\": {\n%s\n\t},", k, v)
		i++
	}
	full := fmt.Sprintf("package %s\n\nvar Values = map[string]map[string]string{\n%s\n}", l, strings.Join(vars[:], "\n"))
	f := fmt.Sprintf("dictionary/%s/vars.go", l)
	file, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.WriteString(full)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func processVars(l string, s string, v map[string]string) string {
	fmt.Printf("Process %s:%s\n", l, s)
	//fmt.Println(v)
	d := fmt.Sprintf("dictionary/%s", l)
	f := fmt.Sprintf("%s/%s.yaml", d, s)
	if _, err := os.Stat(d); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(d, 0755); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	var m map[string]string
	if _, err := os.Stat(f); errors.Is(err, os.ErrNotExist) {

		m = map[string]string{}
	} else {
		y, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = yaml.Unmarshal(y, &m)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	for k, j := range v {
		_, ok := m[k]
		if !ok {
			m[k] = j
		}
	}

	file, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	ym, err := yaml.Marshal(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = file.Write(ym)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	vars := make([]string, len(m))
	i := 0
	for k, j := range m {
		vars[i] = fmt.Sprintf("\t\t\"%s\": \"%s\", ", k, j)
		i++
	}
	return strings.Join(vars[:], "\n")
}

func init() {
	rootCmd.AddCommand(dictionaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dictionaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dictionaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
