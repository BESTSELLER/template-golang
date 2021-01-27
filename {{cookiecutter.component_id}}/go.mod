module github.com/{{cookiecutter.destination.organization}}/{{cookiecutter.component_id}}

go 1.15

require (
	{% if cookiecutter.use_logrus_logging == "y" -%}github.com/sirupsen/logrus v1.7.0{%- endif %}
	{% if cookiecutter.use_cobra_cmd == "y" -%}github.com/spf13/cobra v0.0.7{%- endif %}
	{% if cookiecutter.use_viper_config == "y" -%}github.com/spf13/viper v1.7.1{%- endif %}
)
