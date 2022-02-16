package chaos

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceChaosApplyFailure() *schema.Resource {
	return &schema.Resource{
		Description: "Simple resource that fails on apply",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Name of the collection.",
				Required:    true,
			},
		},
		ReadContext:   resourceChaosApplyFailureRead,
		CreateContext: resourceChaosApplyFailureCreate,
		UpdateContext: resourceChaosApplyFailureUpdate,
		DeleteContext: resourceChaosApplyFailureDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceChaosApplyFailureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	id := d.Id()
	if err := d.Set("name", id); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceChaosApplyFailureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	d.SetId(name)

	return diag.FromErr(errors.New("scheduled failure"))
}

func resourceChaosApplyFailureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceChaosApplyFailureRead(ctx, d, meta)
}

func resourceChaosApplyFailureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.SetId("")
	return diags
}
