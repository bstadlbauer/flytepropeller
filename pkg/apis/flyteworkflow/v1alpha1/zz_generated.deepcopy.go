// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alias.
func (in *Alias) DeepCopy() *Alias {
	if in == nil {
		return nil
	}
	out := new(Alias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Binding.
func (in *Binding) DeepCopy() *Binding {
	if in == nil {
		return nil
	}
	out := new(Binding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BooleanExpression.
func (in *BooleanExpression) DeepCopy() *BooleanExpression {
	if in == nil {
		return nil
	}
	out := new(BooleanExpression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BranchNodeSpec) DeepCopyInto(out *BranchNodeSpec) {
	*out = *in
	in.If.DeepCopyInto(&out.If)
	if in.ElseIf != nil {
		in, out := &in.ElseIf, &out.ElseIf
		*out = make([]*IfBlock, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(IfBlock)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Else != nil {
		in, out := &in.Else, &out.Else
		*out = new(string)
		**out = **in
	}
	if in.ElseFail != nil {
		in, out := &in.ElseFail, &out.ElseFail
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BranchNodeSpec.
func (in *BranchNodeSpec) DeepCopy() *BranchNodeSpec {
	if in == nil {
		return nil
	}
	out := new(BranchNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BranchNodeStatus) DeepCopyInto(out *BranchNodeStatus) {
	*out = *in
	out.MutableStruct = in.MutableStruct
	if in.FinalizedNodeID != nil {
		in, out := &in.FinalizedNodeID, &out.FinalizedNodeID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BranchNodeStatus.
func (in *BranchNodeStatus) DeepCopy() *BranchNodeStatus {
	if in == nil {
		return nil
	}
	out := new(BranchNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Connections.
func (in *Connections) DeepCopy() *Connections {
	if in == nil {
		return nil
	}
	out := new(Connections)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicNodeStatus) DeepCopyInto(out *DynamicNodeStatus) {
	*out = *in
	out.MutableStruct = in.MutableStruct
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicNodeStatus.
func (in *DynamicNodeStatus) DeepCopy() *DynamicNodeStatus {
	if in == nil {
		return nil
	}
	out := new(DynamicNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Error.
func (in *Error) DeepCopy() *Error {
	if in == nil {
		return nil
	}
	out := new(Error)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutionConfig) DeepCopyInto(out *ExecutionConfig) {
	*out = *in
	if in.TaskPluginImpls != nil {
		in, out := &in.TaskPluginImpls, &out.TaskPluginImpls
		*out = make(map[string]TaskPluginOverride, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	out.MaxParallelism = in.MaxParallelism
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutionConfig.
func (in *ExecutionConfig) DeepCopy() *ExecutionConfig {
	if in == nil {
		return nil
	}
	out := new(ExecutionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutionError.
func (in *ExecutionError) DeepCopy() *ExecutionError {
	if in == nil {
		return nil
	}
	out := new(ExecutionError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlyteWorkflow) DeepCopyInto(out *FlyteWorkflow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.WorkflowSpec != nil {
		in, out := &in.WorkflowSpec, &out.WorkflowSpec
		*out = new(WorkflowSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkflowMeta != nil {
		in, out := &in.WorkflowMeta, &out.WorkflowMeta
		*out = new(WorkflowMeta)
		**out = **in
	}
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = (*in).DeepCopy()
	}
	in.ExecutionID.DeepCopyInto(&out.ExecutionID)
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make(map[string]*TaskSpec, len(*in))
		for key, val := range *in {
			var outVal *TaskSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = (*in).DeepCopy()
			}
			(*out)[key] = outVal
		}
	}
	if in.SubWorkflows != nil {
		in, out := &in.SubWorkflows, &out.SubWorkflows
		*out = make(map[string]*WorkflowSpec, len(*in))
		for key, val := range *in {
			var outVal *WorkflowSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(WorkflowSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	out.NodeDefaults = in.NodeDefaults
	if in.AcceptedAt != nil {
		in, out := &in.AcceptedAt, &out.AcceptedAt
		*out = (*in).DeepCopy()
	}
	out.SecurityContext = in.SecurityContext
	in.Status.DeepCopyInto(&out.Status)
	in.RawOutputDataConfig.DeepCopyInto(&out.RawOutputDataConfig)
	in.ExecutionConfig.DeepCopyInto(&out.ExecutionConfig)
	if in.DataReferenceConstructor != nil {
		out.DataReferenceConstructor = in.DataReferenceConstructor
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlyteWorkflow.
func (in *FlyteWorkflow) DeepCopy() *FlyteWorkflow {
	if in == nil {
		return nil
	}
	out := new(FlyteWorkflow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlyteWorkflow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlyteWorkflowList) DeepCopyInto(out *FlyteWorkflowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlyteWorkflow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlyteWorkflowList.
func (in *FlyteWorkflowList) DeepCopy() *FlyteWorkflowList {
	if in == nil {
		return nil
	}
	out := new(FlyteWorkflowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlyteWorkflowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identifier.
func (in *Identifier) DeepCopy() *Identifier {
	if in == nil {
		return nil
	}
	out := new(Identifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IfBlock) DeepCopyInto(out *IfBlock) {
	*out = *in
	in.Condition.DeepCopyInto(&out.Condition)
	if in.ThenNode != nil {
		in, out := &in.ThenNode, &out.ThenNode
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IfBlock.
func (in *IfBlock) DeepCopy() *IfBlock {
	if in == nil {
		return nil
	}
	out := new(IfBlock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Inputs.
func (in *Inputs) DeepCopy() *Inputs {
	if in == nil {
		return nil
	}
	out := new(Inputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutableStruct) DeepCopyInto(out *MutableStruct) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutableStruct.
func (in *MutableStruct) DeepCopy() *MutableStruct {
	if in == nil {
		return nil
	}
	out := new(MutableStruct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeDefaults) DeepCopyInto(out *NodeDefaults) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeDefaults.
func (in *NodeDefaults) DeepCopy() *NodeDefaults {
	if in == nil {
		return nil
	}
	out := new(NodeDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMetadata.
func (in *NodeMetadata) DeepCopy() *NodeMetadata {
	if in == nil {
		return nil
	}
	out := new(NodeMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSpec) DeepCopyInto(out *NodeSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.BranchNode != nil {
		in, out := &in.BranchNode, &out.BranchNode
		*out = new(BranchNodeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.TaskRef != nil {
		in, out := &in.TaskRef, &out.TaskRef
		*out = new(string)
		**out = **in
	}
	if in.WorkflowNode != nil {
		in, out := &in.WorkflowNode, &out.WorkflowNode
		*out = new(WorkflowNodeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.InputBindings != nil {
		in, out := &in.InputBindings, &out.InputBindings
		*out = make([]*Binding, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(v1.ConfigMap)
		(*in).DeepCopyInto(*out)
	}
	if in.RetryStrategy != nil {
		in, out := &in.RetryStrategy, &out.RetryStrategy
		*out = new(RetryStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.OutputAliases != nil {
		in, out := &in.OutputAliases, &out.OutputAliases
		*out = make([]Alias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExecutionDeadline != nil {
		in, out := &in.ExecutionDeadline, &out.ExecutionDeadline
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.ActiveDeadline != nil {
		in, out := &in.ActiveDeadline, &out.ActiveDeadline
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.Interruptibe != nil {
		in, out := &in.Interruptibe, &out.Interruptibe
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSpec.
func (in *NodeSpec) DeepCopy() *NodeSpec {
	if in == nil {
		return nil
	}
	out := new(NodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
	out.MutableStruct = in.MutableStruct
	if in.QueuedAt != nil {
		in, out := &in.QueuedAt, &out.QueuedAt
		*out = (*in).DeepCopy()
	}
	if in.StartedAt != nil {
		in, out := &in.StartedAt, &out.StartedAt
		*out = (*in).DeepCopy()
	}
	if in.StoppedAt != nil {
		in, out := &in.StoppedAt, &out.StoppedAt
		*out = (*in).DeepCopy()
	}
	if in.LastUpdatedAt != nil {
		in, out := &in.LastUpdatedAt, &out.LastUpdatedAt
		*out = (*in).DeepCopy()
	}
	if in.LastAttemptStartedAt != nil {
		in, out := &in.LastAttemptStartedAt, &out.LastAttemptStartedAt
		*out = (*in).DeepCopy()
	}
	if in.ParentNode != nil {
		in, out := &in.ParentNode, &out.ParentNode
		*out = new(string)
		**out = **in
	}
	if in.ParentTask != nil {
		in, out := &in.ParentTask, &out.ParentTask
		*out = (*in).DeepCopy()
	}
	if in.BranchStatus != nil {
		in, out := &in.BranchStatus, &out.BranchStatus
		*out = new(BranchNodeStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.SubNodeStatus != nil {
		in, out := &in.SubNodeStatus, &out.SubNodeStatus
		*out = make(map[string]*NodeStatus, len(*in))
		for key, val := range *in {
			var outVal *NodeStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(NodeStatus)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.WorkflowNodeStatus != nil {
		in, out := &in.WorkflowNodeStatus, &out.WorkflowNodeStatus
		*out = new(WorkflowNodeStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.TaskNodeStatus != nil {
		in, out := &in.TaskNodeStatus, &out.TaskNodeStatus
		*out = (*in).DeepCopy()
	}
	if in.DynamicNodeStatus != nil {
		in, out := &in.DynamicNodeStatus, &out.DynamicNodeStatus
		*out = new(DynamicNodeStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = (*in).DeepCopy()
	}
	if in.DataReferenceConstructor != nil {
		out.DataReferenceConstructor = in.DataReferenceConstructor
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputVarMap.
func (in *OutputVarMap) DeepCopy() *OutputVarMap {
	if in == nil {
		return nil
	}
	out := new(OutputVarMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RawOutputDataConfig.
func (in *RawOutputDataConfig) DeepCopy() *RawOutputDataConfig {
	if in == nil {
		return nil
	}
	out := new(RawOutputDataConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryStrategy) DeepCopyInto(out *RetryStrategy) {
	*out = *in
	if in.MinAttempts != nil {
		in, out := &in.MinAttempts, &out.MinAttempts
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryStrategy.
func (in *RetryStrategy) DeepCopy() *RetryStrategy {
	if in == nil {
		return nil
	}
	out := new(RetryStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskExecutionIdentifier.
func (in *TaskExecutionIdentifier) DeepCopy() *TaskExecutionIdentifier {
	if in == nil {
		return nil
	}
	out := new(TaskExecutionIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskPluginOverride) DeepCopyInto(out *TaskPluginOverride) {
	*out = *in
	if in.PluginIDs != nil {
		in, out := &in.PluginIDs, &out.PluginIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskPluginOverride.
func (in *TaskPluginOverride) DeepCopy() *TaskPluginOverride {
	if in == nil {
		return nil
	}
	out := new(TaskPluginOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskSpec.
func (in *TaskSpec) DeepCopy() *TaskSpec {
	if in == nil {
		return nil
	}
	out := new(TaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowExecutionIdentifier.
func (in *WorkflowExecutionIdentifier) DeepCopy() *WorkflowExecutionIdentifier {
	if in == nil {
		return nil
	}
	out := new(WorkflowExecutionIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowMeta) DeepCopyInto(out *WorkflowMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowMeta.
func (in *WorkflowMeta) DeepCopy() *WorkflowMeta {
	if in == nil {
		return nil
	}
	out := new(WorkflowMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowNodeSpec) DeepCopyInto(out *WorkflowNodeSpec) {
	*out = *in
	if in.LaunchPlanRefID != nil {
		in, out := &in.LaunchPlanRefID, &out.LaunchPlanRefID
		*out = (*in).DeepCopy()
	}
	if in.SubWorkflowReference != nil {
		in, out := &in.SubWorkflowReference, &out.SubWorkflowReference
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowNodeSpec.
func (in *WorkflowNodeSpec) DeepCopy() *WorkflowNodeSpec {
	if in == nil {
		return nil
	}
	out := new(WorkflowNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowNodeStatus) DeepCopyInto(out *WorkflowNodeStatus) {
	*out = *in
	out.MutableStruct = in.MutableStruct
	if in.ExecutionError != nil {
		in, out := &in.ExecutionError, &out.ExecutionError
		*out = new(core.ExecutionError)
		*out = *in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowNodeStatus.
func (in *WorkflowNodeStatus) DeepCopy() *WorkflowNodeStatus {
	if in == nil {
		return nil
	}
	out := new(WorkflowNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowSpec) DeepCopyInto(out *WorkflowSpec) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[string]*NodeSpec, len(*in))
		for key, val := range *in {
			var outVal *NodeSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(NodeSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	in.DeprecatedConnections.DeepCopyInto(&out.DeprecatedConnections)
	if in.OnFailure != nil {
		in, out := &in.OnFailure, &out.OnFailure
		*out = new(NodeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = (*in).DeepCopy()
	}
	if in.OutputBindings != nil {
		in, out := &in.OutputBindings, &out.OutputBindings
		*out = make([]*Binding, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowSpec.
func (in *WorkflowSpec) DeepCopy() *WorkflowSpec {
	if in == nil {
		return nil
	}
	out := new(WorkflowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowStatus) DeepCopyInto(out *WorkflowStatus) {
	*out = *in
	if in.StartedAt != nil {
		in, out := &in.StartedAt, &out.StartedAt
		*out = (*in).DeepCopy()
	}
	if in.StoppedAt != nil {
		in, out := &in.StoppedAt, &out.StoppedAt
		*out = (*in).DeepCopy()
	}
	if in.LastUpdatedAt != nil {
		in, out := &in.LastUpdatedAt, &out.LastUpdatedAt
		*out = (*in).DeepCopy()
	}
	if in.NodeStatus != nil {
		in, out := &in.NodeStatus, &out.NodeStatus
		*out = make(map[string]*NodeStatus, len(*in))
		for key, val := range *in {
			var outVal *NodeStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(NodeStatus)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = (*in).DeepCopy()
	}
	if in.DataReferenceConstructor != nil {
		out.DataReferenceConstructor = in.DataReferenceConstructor
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowStatus.
func (in *WorkflowStatus) DeepCopy() *WorkflowStatus {
	if in == nil {
		return nil
	}
	out := new(WorkflowStatus)
	in.DeepCopyInto(out)
	return out
}
