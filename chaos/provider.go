package chaos

import (
	"context"
	"math/rand"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},

		ResourcesMap: map[string]*schema.Resource{
			"chaos_apply_failure":        resourceChaosApplyFailure(),
			"chaos_apply_random_failure": resourceChaosApplyRandomFailure(),
			"chaos_plan_failure":         resourceChaosPlanFailure(),
			"chaos_plan_random_failure":  resourceChaosPlanFailure(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	rand.Seed(time.Now().UnixNano())

	return nil, nil
}
