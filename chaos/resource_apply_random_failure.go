package chaos

import (
	"context"
	"errors"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceChaosApplyRandomFailure() *schema.Resource {
	return &schema.Resource{
		Description: "Randomly fails on apply",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Name of the chaos resource",
				Required:    true,
			},
			"failure_rate": {
				Type:         schema.TypeFloat,
				Description:  "Rate of the failure",
				Default:      1.0,
				ValidateFunc: validation.FloatBetween(0, 1.0),
				Optional:     true,
			},
		},
		ReadContext:   resourceChaosApplyRandomFailureRead,
		CreateContext: resourceChaosApplyRandomFailureCreate,
		UpdateContext: resourceChaosApplyRandomFailureUpdate,
		DeleteContext: resourceChaosApplyRandomFailureDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceChaosApplyRandomFailureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	id := d.Id()
	if err := d.Set("name", id); err != nil {
		return diag.FromErr(err)
	}

	rate := d.Get("failure_rate").(float64)
	if err := d.Set("failure_rate", rate); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceChaosApplyRandomFailureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	n := rand.Float64()
	rate := d.Get("failure_rate").(float64)

	if rate >= n {
		diag.FromErr(errors.New("scheduled failure"))
	}

	d.SetId(name)
	return resourceChaosApplyRandomFailureRead(ctx, d, meta)
}

func resourceChaosApplyRandomFailureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceChaosApplyRandomFailureRead(ctx, d, meta)
}

func resourceChaosApplyRandomFailureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.SetId("")
	return diags
}
