// This file contains the implementation of Gitlab.

package app

import (
	"fmt"
	"github.com/Bitspark/go-bitnode/bitnode"
	"gitlab/util"
	"reflect"
)

// Struct definition for Gitlab.

// Gitlab is the main sparkable.
type Gitlab struct {
	bitnode.System
}

// Gitlab methods.

// ConnectGitlabInstance description: Method connectGitlabInstance of Gitlab.
func (s *Gitlab) ConnectGitlabInstance() error {
	// TODO: Implement method.
	return fmt.Errorf("method connectGitlabInstance not implemented yet")
}

// CreateIssue description: Method createIssue of Gitlab.
func (s *Gitlab) CreateIssue(issue Issue) error {
	// TODO: Implement method.
	return fmt.Errorf("method createIssue not implemented yet")
}

// AddNoteToIssue description: Method addNoteToIssue of Gitlab.
func (s *Gitlab) AddNoteToIssue(note string, issue Issue) error {
	// TODO: Implement method.
	return fmt.Errorf("method addNoteToIssue not implemented yet")
}

// AddNoteToMergeRequest description: Method addNoteToMergeRequest of Gitlab.
func (s *Gitlab) AddNoteToMergeRequest(note string, mergeRequest MergeRequest) error {
	// TODO: Implement method.
	return fmt.Errorf("method addNoteToMergeRequest not implemented yet")
}

// CreateNewProject description: Method createNewProject of Gitlab.
func (s *Gitlab) CreateNewProject(project Project) error {
	// TODO: Implement method.
	return fmt.Errorf("method createNewProject not implemented yet")
}

// Lifecycle callbacks.

// lifecycleCreate is called when the container has been created.
func (s *Gitlab) lifecycleCreate(vals ...bitnode.HubItem) error {
	// TODO: Add startup logic here which is called when the spark is created.
	return nil
}

// lifecycleLoad is called when the container has been started (after lifecycleCreate) or restarted.
func (s *Gitlab) lifecycleLoad(vals ...bitnode.HubItem) error {
	// TODO: Add startup logic here which is called after the spark has been created.

	s.SetMessage("Gitlab running...")
	s.SetStatus(bitnode.SystemStatusRunning)

	return nil
}

// DO NOT CHANGE THE FOLLOWING CODE UNLESS YOU KNOW WHAT YOU ARE DOING.

func (s *Gitlab) Update(values ...string) error {
	sv := reflect.ValueOf(*s)
	st := reflect.TypeOf(*s)
	if len(values) == 0 {
		for i := 0; i < st.NumField(); i++ {
			values = append(values, st.Field(i).Name)
		}
	}
	for _, value := range values {
		ft, ok := st.FieldByName(value)
		if !ok {
			return fmt.Errorf("field '%s' not found in %s", value, st.Name())
		}
		fv := sv.FieldByName(value)
		if !fv.IsValid() {
			return fmt.Errorf("field '%s' not found in %s", value, st.Name())
		}
		val, err := util.InterfaceFromValue(fv.Interface())
		if err != nil {
			return err
		}
		hubName := ft.Tag.Get("json")
		if err := s.GetHub(hubName).Set("", val); err != nil {
			return err
		}
	}
	return nil
}

// Init attaches the methods of the Gitlab to the respective handlers.
func (s *Gitlab) Init() error {
	// METHODS

	s.GetHub("connectGitlabInstance").Handle(bitnode.NewNativeFunction(func(creds bitnode.Credentials, vals ...bitnode.HubItem) ([]bitnode.HubItem, error) {
		err := s.ConnectGitlabInstance()
		if err != nil {
			return nil, err
		}

		return []bitnode.HubItem{}, nil
	}))

	s.GetHub("createIssue").Handle(bitnode.NewNativeFunction(func(creds bitnode.Credentials, vals ...bitnode.HubItem) ([]bitnode.HubItem, error) {
		issue, err := util.ValueFromInterface[Issue](vals[0])
		if err != nil {
			return nil, err
		}

		err = s.CreateIssue(issue)
		if err != nil {
			return nil, err
		}

		return []bitnode.HubItem{}, nil
	}))

	s.GetHub("addNoteToIssue").Handle(bitnode.NewNativeFunction(func(creds bitnode.Credentials, vals ...bitnode.HubItem) ([]bitnode.HubItem, error) {
		note, err := util.ValueFromInterface[string](vals[0])
		if err != nil {
			return nil, err
		}
		issue, err := util.ValueFromInterface[Issue](vals[1])
		if err != nil {
			return nil, err
		}

		err = s.AddNoteToIssue(note, issue)
		if err != nil {
			return nil, err
		}

		return []bitnode.HubItem{}, nil
	}))

	s.GetHub("addNoteToMergeRequest").Handle(bitnode.NewNativeFunction(func(creds bitnode.Credentials, vals ...bitnode.HubItem) ([]bitnode.HubItem, error) {
		note, err := util.ValueFromInterface[string](vals[0])
		if err != nil {
			return nil, err
		}
		mergeRequest, err := util.ValueFromInterface[MergeRequest](vals[1])
		if err != nil {
			return nil, err
		}

		err = s.AddNoteToMergeRequest(note, mergeRequest)
		if err != nil {
			return nil, err
		}

		return []bitnode.HubItem{}, nil
	}))

	s.GetHub("createNewProject").Handle(bitnode.NewNativeFunction(func(creds bitnode.Credentials, vals ...bitnode.HubItem) ([]bitnode.HubItem, error) {
		project, err := util.ValueFromInterface[Project](vals[0])
		if err != nil {
			return nil, err
		}

		err = s.CreateNewProject(project)
		if err != nil {
			return nil, err
		}

		return []bitnode.HubItem{}, nil
	}))

	// VALUES

	// CHANNELS

	// LIFECYCLE EVENTS

	s.AddCallback(bitnode.LifecycleCreate, bitnode.NewNativeEvent(func(vals ...bitnode.HubItem) error {
		return s.lifecycleCreate(vals...)
	}))

	s.AddCallback(bitnode.LifecycleLoad, bitnode.NewNativeEvent(func(vals ...bitnode.HubItem) error {
		return s.lifecycleLoad(vals...)
	}))

	return nil
}
