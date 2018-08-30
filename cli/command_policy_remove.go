package cli

import (
	"context"

	"github.com/kopia/kopia/policy"
	"github.com/kopia/kopia/repo"
)

var (
	policyRemoveCommand = policyCommands.Command("remove", "Remove snapshot policy for a single directory, user@host or a global policy.").Alias("rm").Alias("delete")
	policyRemoveTargets = policyRemoveCommand.Arg("target", "Target of a policy ('global','user@host','@host') or a path").Strings()
	policyRemoveGlobal  = policyRemoveCommand.Flag("global", "Set global policy").Bool()
)

func init() {
	policyRemoveCommand.Action(repositoryAction(removePolicy))
}

func removePolicy(ctx context.Context, rep *repo.Repository) error {
	mgr := policy.NewPolicyManager(rep)

	targets, err := policyTargets(mgr, policyRemoveGlobal, policyRemoveTargets)
	if err != nil {
		return err
	}

	for _, target := range targets {
		log.Infof("Removing policy on %q...", target)
		if err := mgr.RemovePolicy(target); err != nil {
			return err
		}
	}

	return nil
}
