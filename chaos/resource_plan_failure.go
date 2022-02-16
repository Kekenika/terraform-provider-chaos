package chaos

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceChaosPlanFailure() *schema.Resource {
	return &schema.Resource{
		Description: "Simple resource that fails on plan",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Name of the collection.",
				Required:    true,
			},
		},
		ReadContext:   resourceChaosPlanFailureRead,
		CreateContext: resourceChaosPlanFailureCreate,
		UpdateContext: resourceChaosPlanFailureUpdate,
		DeleteContext: resourceChaosPlanFailureDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceChaosPlanFailureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.FromErr(errors.New("scheduled failure"))
}

func resourceChaosPlanFailureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	d.SetId(name)
	return resourceChaosPlanFailureRead(ctx, d, meta)
}

func resourceChaosPlanFailureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceChaosPlanFailureRead(ctx, d, meta)
}

func resourceChaosPlanFailureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.SetId("")
	return diags
}
