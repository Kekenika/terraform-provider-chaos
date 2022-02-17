package chaos

import (
	"context"
	"errors"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceChaosPlanRandomFailure() *schema.Resource {
	return &schema.Resource{
		Description: "Randomly fails on plan",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Name of the chaos resource",
				Required:    true,
			},
			"failure_rate": {
				Type:         schema.TypeInt,
				Description:  "Rate of the failure",
				Default:      1.0,
				ValidateFunc: validation.FloatBetween(0, 1.0),
				Optional:     true,
			},
		},
		ReadContext:   resourceChaosPlanRandomFailureRead,
		CreateContext: resourceChaosPlanRandomFailureCreate,
		UpdateContext: resourceChaosPlanRandomFailureUpdate,
		DeleteContext: resourceChaosPlanRandomFailureDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceChaosPlanRandomFailureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	id := d.Id()
	if err := d.Set("name", id); err != nil {
		return diag.FromErr(err)
	}

	rate := d.Get("failure_rate").(float64)
	n := rand.Float64()

	if rate >= n {
		diag.FromErr(errors.New("scheduled failure"))
	}
	if err := d.Set("failure_rate", rate); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceChaosPlanRandomFailureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	d.SetId(name)
	return resourceChaosPlanRandomFailureRead(ctx, d, meta)
}

func resourceChaosPlanRandomFailureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceChaosPlanRandomFailureRead(ctx, d, meta)
}

func resourceChaosPlanRandomFailureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.SetId("")
	return diags
}
