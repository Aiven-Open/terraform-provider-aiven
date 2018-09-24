package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jelmersnoeck/aiven"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,

		Schema: map[string]*schema.Schema{
			"card_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Credit card ID",
			},
			"cloud": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Target cloud",
			},
			"project": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Project name",
			},
			"ca_crt": {
				Type:        schema.TypeString,
				Sensitive:   true,
				Computed:    true,
				Description: "Project CA Certificate",
			},
		},
	}
}

func dataSourceProject() *schema.Resource {
	return &schema.Resource{
		Read: resourceProjectRead,

		Schema: map[string]*schema.Schema{
			"card_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Credit card ID",
			},
			"cloud": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Target cloud",
			},
			"project": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Project name",
			},
			"ca_crt": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Project CA Certificate",
			},
		},
	}
}

func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*aiven.Client)
	project, err := client.Projects.Create(
		aiven.CreateProjectRequest{
			CardID:  d.Get("card_id").(string),
			Cloud:   d.Get("cloud").(string),
			Project: d.Get("project").(string),
		},
	)
	if err != nil {
		return err
	}

	d.SetId(project.Name + "!")
	return nil
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*aiven.Client)

	project, err := client.Projects.Get(d.Get("project").(string))
	if err != nil {
		return err
	}
	ca, err := client.CA.Get(d.Get("project").(string))
	if err != nil {
		return err
	}

	d.Set("project", project.Name)
	d.Set("ca_crt", ca)
	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*aiven.Client)

	project, err := client.Projects.Update(
		d.Get("project").(string),
		aiven.UpdateProjectRequest{
			CardID: d.Get("card_id").(string),
			Cloud:  d.Get("cloud").(string),
		},
	)
	if err != nil {
		return err
	}

	d.SetId(project.Name + "!")
	return nil
}

func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*aiven.Client)

	return client.Projects.Delete(d.Get("project").(string))
}
