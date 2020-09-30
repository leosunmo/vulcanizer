module github.com/leosunmo/vulcanizer

go 1.15

require (
	github.com/github/vulcanizer v0.5.1
	github.com/jeremywohl/flatten v0.0.0-20180616191615-6ba45bff2452
	github.com/olekukonko/tablewriter v0.0.0-20180506121414-d4647c9c7a84
	github.com/parnurzeal/gorequest v0.2.15
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.3.2
	github.com/tidwall/gjson v1.1.3
	gotest.tools v2.2.0+incompatible
)

replace github.com/github/vulcanizer v0.5.1 => github.com/leosunmo/vulcanizer v0.5.3
