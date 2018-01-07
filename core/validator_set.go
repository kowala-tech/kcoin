package core

type ValidatorSet struct {
	Validators []*Validator
	Proposer   *Validator
}

func NewValidatorSet(validators []*Validator) *ValidatorSet {
	//validators := make([]*Validator, len(validators))

	/*
		for i, val := range vals {
			validators[i] = val.Copy()
		}

		// @TODO (rgeraldes) - create a sorter for the validator type
		//sort.Sort(ValidatorsByAddress(validators))
	*/

	vs := &ValidatorSet{
		Validators: validators,
	}

	/*
		if vals != nil {
			vs.IncrementAccum(1)
		}
	*/

	return vs
}

func (set *ValidatorSet) Size() int { return len(set.Validators) }
