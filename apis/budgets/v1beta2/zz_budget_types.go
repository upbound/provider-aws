// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AutoAdjustDataInitParameters struct {
	AutoAdjustType *string `json:"autoAdjustType,omitempty" tf:"auto_adjust_type,omitempty"`

	HistoricalOptions []HistoricalOptionsInitParameters `json:"historicalOptions,omitempty" tf:"historical_options,omitempty"`
}

type AutoAdjustDataObservation struct {
	AutoAdjustType *string `json:"autoAdjustType,omitempty" tf:"auto_adjust_type,omitempty"`

	HistoricalOptions []HistoricalOptionsObservation `json:"historicalOptions,omitempty" tf:"historical_options,omitempty"`

	LastAutoAdjustTime *string `json:"lastAutoAdjustTime,omitempty" tf:"last_auto_adjust_time,omitempty"`
}

type AutoAdjustDataParameters struct {

	// +kubebuilder:validation:Optional
	AutoAdjustType *string `json:"autoAdjustType" tf:"auto_adjust_type,omitempty"`

	// +kubebuilder:validation:Optional
	HistoricalOptions []HistoricalOptionsParameters `json:"historicalOptions,omitempty" tf:"historical_options,omitempty"`
}

type BudgetInitParameters struct {

	// The ID of the target account for budget. Will use current user's account_id by default if omitted.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Object containing [AutoAdjustData] which determines the budget amount for an auto-adjusting budget.
	AutoAdjustData []AutoAdjustDataInitParameters `json:"autoAdjustData,omitempty" tf:"auto_adjust_data,omitempty"`

	// Whether this budget tracks monetary cost or usage.
	BudgetType *string `json:"budgetType,omitempty" tf:"budget_type,omitempty"`

	// A list of CostFilter name/values pair to apply to budget.
	CostFilter []CostFilterInitParameters `json:"costFilter,omitempty" tf:"cost_filter,omitempty"`

	// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions.
	CostTypes []CostTypesInitParameters `json:"costTypes,omitempty" tf:"cost_types,omitempty"`

	// The amount of cost or usage being measured for a budget.
	LimitAmount *string `json:"limitAmount,omitempty" tf:"limit_amount,omitempty"`

	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See Spend documentation.
	LimitUnit *string `json:"limitUnit,omitempty" tf:"limit_unit,omitempty"`

	// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification.
	Notification []NotificationInitParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// Object containing Planned Budget Limits. Can be used multiple times to plan more than one budget limit. See PlannedBudgetLimits documentation.
	PlannedLimit []PlannedLimitInitParameters `json:"plannedLimit,omitempty" tf:"planned_limit,omitempty"`

	// The end of the time period covered by the budget. There are no restrictions on the end date. Format: 2017-01-01_12:00.
	TimePeriodEnd *string `json:"timePeriodEnd,omitempty" tf:"time_period_end,omitempty"`

	// The start of the time period covered by the budget. If you don't specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: 2017-01-01_12:00.
	TimePeriodStart *string `json:"timePeriodStart,omitempty" tf:"time_period_start,omitempty"`

	// The length of time until a budget resets the actual and forecasted spend. Valid values: MONTHLY, QUARTERLY, ANNUALLY, and DAILY.
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type BudgetObservation struct {

	// The ID of the target account for budget. Will use current user's account_id by default if omitted.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The ARN of the budget.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Object containing [AutoAdjustData] which determines the budget amount for an auto-adjusting budget.
	AutoAdjustData []AutoAdjustDataObservation `json:"autoAdjustData,omitempty" tf:"auto_adjust_data,omitempty"`

	// Whether this budget tracks monetary cost or usage.
	BudgetType *string `json:"budgetType,omitempty" tf:"budget_type,omitempty"`

	// A list of CostFilter name/values pair to apply to budget.
	CostFilter []CostFilterObservation `json:"costFilter,omitempty" tf:"cost_filter,omitempty"`

	// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions.
	CostTypes []CostTypesObservation `json:"costTypes,omitempty" tf:"cost_types,omitempty"`

	// id of resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The amount of cost or usage being measured for a budget.
	LimitAmount *string `json:"limitAmount,omitempty" tf:"limit_amount,omitempty"`

	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See Spend documentation.
	LimitUnit *string `json:"limitUnit,omitempty" tf:"limit_unit,omitempty"`

	// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification.
	Notification []NotificationObservation `json:"notification,omitempty" tf:"notification,omitempty"`

	// Object containing Planned Budget Limits. Can be used multiple times to plan more than one budget limit. See PlannedBudgetLimits documentation.
	PlannedLimit []PlannedLimitObservation `json:"plannedLimit,omitempty" tf:"planned_limit,omitempty"`

	// The end of the time period covered by the budget. There are no restrictions on the end date. Format: 2017-01-01_12:00.
	TimePeriodEnd *string `json:"timePeriodEnd,omitempty" tf:"time_period_end,omitempty"`

	// The start of the time period covered by the budget. If you don't specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: 2017-01-01_12:00.
	TimePeriodStart *string `json:"timePeriodStart,omitempty" tf:"time_period_start,omitempty"`

	// The length of time until a budget resets the actual and forecasted spend. Valid values: MONTHLY, QUARTERLY, ANNUALLY, and DAILY.
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type BudgetParameters struct {

	// The ID of the target account for budget. Will use current user's account_id by default if omitted.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Object containing [AutoAdjustData] which determines the budget amount for an auto-adjusting budget.
	// +kubebuilder:validation:Optional
	AutoAdjustData []AutoAdjustDataParameters `json:"autoAdjustData,omitempty" tf:"auto_adjust_data,omitempty"`

	// Whether this budget tracks monetary cost or usage.
	// +kubebuilder:validation:Optional
	BudgetType *string `json:"budgetType,omitempty" tf:"budget_type,omitempty"`

	// A list of CostFilter name/values pair to apply to budget.
	// +kubebuilder:validation:Optional
	CostFilter []CostFilterParameters `json:"costFilter,omitempty" tf:"cost_filter,omitempty"`

	// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions.
	// +kubebuilder:validation:Optional
	CostTypes []CostTypesParameters `json:"costTypes,omitempty" tf:"cost_types,omitempty"`

	// The amount of cost or usage being measured for a budget.
	// +kubebuilder:validation:Optional
	LimitAmount *string `json:"limitAmount,omitempty" tf:"limit_amount,omitempty"`

	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See Spend documentation.
	// +kubebuilder:validation:Optional
	LimitUnit *string `json:"limitUnit,omitempty" tf:"limit_unit,omitempty"`

	// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification.
	// +kubebuilder:validation:Optional
	Notification []NotificationParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// Object containing Planned Budget Limits. Can be used multiple times to plan more than one budget limit. See PlannedBudgetLimits documentation.
	// +kubebuilder:validation:Optional
	PlannedLimit []PlannedLimitParameters `json:"plannedLimit,omitempty" tf:"planned_limit,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The end of the time period covered by the budget. There are no restrictions on the end date. Format: 2017-01-01_12:00.
	// +kubebuilder:validation:Optional
	TimePeriodEnd *string `json:"timePeriodEnd,omitempty" tf:"time_period_end,omitempty"`

	// The start of the time period covered by the budget. If you don't specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: 2017-01-01_12:00.
	// +kubebuilder:validation:Optional
	TimePeriodStart *string `json:"timePeriodStart,omitempty" tf:"time_period_start,omitempty"`

	// The length of time until a budget resets the actual and forecasted spend. Valid values: MONTHLY, QUARTERLY, ANNUALLY, and DAILY.
	// +kubebuilder:validation:Optional
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type CostFilterInitParameters struct {

	// The name of a budget. Unique within accounts.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type CostFilterObservation struct {

	// The name of a budget. Unique within accounts.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type CostFilterParameters struct {

	// The name of a budget. Unique within accounts.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type CostTypesInitParameters struct {

	// A boolean value whether to include credits in the cost budget. Defaults to true
	IncludeCredit *bool `json:"includeCredit,omitempty" tf:"include_credit,omitempty"`

	// Whether a budget includes discounts. Defaults to true
	IncludeDiscount *bool `json:"includeDiscount,omitempty" tf:"include_discount,omitempty"`

	// A boolean value whether to include other subscription costs in the cost budget. Defaults to true
	IncludeOtherSubscription *bool `json:"includeOtherSubscription,omitempty" tf:"include_other_subscription,omitempty"`

	// A boolean value whether to include recurring costs in the cost budget. Defaults to true
	IncludeRecurring *bool `json:"includeRecurring,omitempty" tf:"include_recurring,omitempty"`

	// A boolean value whether to include refunds in the cost budget. Defaults to true
	IncludeRefund *bool `json:"includeRefund,omitempty" tf:"include_refund,omitempty"`

	// A boolean value whether to include subscriptions in the cost budget. Defaults to true
	IncludeSubscription *bool `json:"includeSubscription,omitempty" tf:"include_subscription,omitempty"`

	// A boolean value whether to include support costs in the cost budget. Defaults to true
	IncludeSupport *bool `json:"includeSupport,omitempty" tf:"include_support,omitempty"`

	// A boolean value whether to include tax in the cost budget. Defaults to true
	IncludeTax *bool `json:"includeTax,omitempty" tf:"include_tax,omitempty"`

	// A boolean value whether to include upfront costs in the cost budget. Defaults to true
	IncludeUpfront *bool `json:"includeUpfront,omitempty" tf:"include_upfront,omitempty"`

	// Whether a budget uses the amortized rate. Defaults to false
	UseAmortized *bool `json:"useAmortized,omitempty" tf:"use_amortized,omitempty"`

	// A boolean value whether to use blended costs in the cost budget. Defaults to false
	UseBlended *bool `json:"useBlended,omitempty" tf:"use_blended,omitempty"`
}

type CostTypesObservation struct {

	// A boolean value whether to include credits in the cost budget. Defaults to true
	IncludeCredit *bool `json:"includeCredit,omitempty" tf:"include_credit,omitempty"`

	// Whether a budget includes discounts. Defaults to true
	IncludeDiscount *bool `json:"includeDiscount,omitempty" tf:"include_discount,omitempty"`

	// A boolean value whether to include other subscription costs in the cost budget. Defaults to true
	IncludeOtherSubscription *bool `json:"includeOtherSubscription,omitempty" tf:"include_other_subscription,omitempty"`

	// A boolean value whether to include recurring costs in the cost budget. Defaults to true
	IncludeRecurring *bool `json:"includeRecurring,omitempty" tf:"include_recurring,omitempty"`

	// A boolean value whether to include refunds in the cost budget. Defaults to true
	IncludeRefund *bool `json:"includeRefund,omitempty" tf:"include_refund,omitempty"`

	// A boolean value whether to include subscriptions in the cost budget. Defaults to true
	IncludeSubscription *bool `json:"includeSubscription,omitempty" tf:"include_subscription,omitempty"`

	// A boolean value whether to include support costs in the cost budget. Defaults to true
	IncludeSupport *bool `json:"includeSupport,omitempty" tf:"include_support,omitempty"`

	// A boolean value whether to include tax in the cost budget. Defaults to true
	IncludeTax *bool `json:"includeTax,omitempty" tf:"include_tax,omitempty"`

	// A boolean value whether to include upfront costs in the cost budget. Defaults to true
	IncludeUpfront *bool `json:"includeUpfront,omitempty" tf:"include_upfront,omitempty"`

	// Whether a budget uses the amortized rate. Defaults to false
	UseAmortized *bool `json:"useAmortized,omitempty" tf:"use_amortized,omitempty"`

	// A boolean value whether to use blended costs in the cost budget. Defaults to false
	UseBlended *bool `json:"useBlended,omitempty" tf:"use_blended,omitempty"`
}

type CostTypesParameters struct {

	// A boolean value whether to include credits in the cost budget. Defaults to true
	// +kubebuilder:validation:Optional
	IncludeCredit *bool `json:"includeCredit,omitempty" tf:"include_credit,omitempty"`

	// Whether a budget includes discounts. Defaults to true
	// +kubebuilder:validation:Optional
	IncludeDiscount *bool `json:"includeDiscount,omitempty" tf:"include_discount,omitempty"`

	// A boolean value whether to include other subscription costs in the cost budget. Defaults to true
	// +kubebuilder:validation:Optional
	IncludeOtherSubscription *bool `json:"includeOtherSubscription,omitempty" tf:"include_other_subscription,omitempty"`

	// A boolean value whether to include recurring costs in the cost budget. Defaults to true
	// +kubebuilder:validation:Optional
	IncludeRecurring *bool `json:"includeRecurring,omitempty" tf:"include_recurring,omitempty"`

	// A boolean value whether to include refunds in the cost budget. Defaults to true
	// +kubebuilder:validation:Optional
	IncludeRefund *bool `json:"includeRefund,omitempty" tf:"include_refund,omitempty"`

	// A boolean value whether to include subscriptions in the cost budget. Defaults to true
	// +kubebuilder:validation:Optional
	IncludeSubscription *bool `json:"includeSubscription,omitempty" tf:"include_subscription,omitempty"`

	// A boolean value whether to include support costs in the cost budget. Defaults to true
	// +kubebuilder:validation:Optional
	IncludeSupport *bool `json:"includeSupport,omitempty" tf:"include_support,omitempty"`

	// A boolean value whether to include tax in the cost budget. Defaults to true
	// +kubebuilder:validation:Optional
	IncludeTax *bool `json:"includeTax,omitempty" tf:"include_tax,omitempty"`

	// A boolean value whether to include upfront costs in the cost budget. Defaults to true
	// +kubebuilder:validation:Optional
	IncludeUpfront *bool `json:"includeUpfront,omitempty" tf:"include_upfront,omitempty"`

	// Whether a budget uses the amortized rate. Defaults to false
	// +kubebuilder:validation:Optional
	UseAmortized *bool `json:"useAmortized,omitempty" tf:"use_amortized,omitempty"`

	// A boolean value whether to use blended costs in the cost budget. Defaults to false
	// +kubebuilder:validation:Optional
	UseBlended *bool `json:"useBlended,omitempty" tf:"use_blended,omitempty"`
}

type HistoricalOptionsInitParameters struct {
	BudgetAdjustmentPeriod *float64 `json:"budgetAdjustmentPeriod,omitempty" tf:"budget_adjustment_period,omitempty"`
}

type HistoricalOptionsObservation struct {
	BudgetAdjustmentPeriod *float64 `json:"budgetAdjustmentPeriod,omitempty" tf:"budget_adjustment_period,omitempty"`

	LookbackAvailablePeriods *float64 `json:"lookbackAvailablePeriods,omitempty" tf:"lookback_available_periods,omitempty"`
}

type HistoricalOptionsParameters struct {

	// +kubebuilder:validation:Optional
	BudgetAdjustmentPeriod *float64 `json:"budgetAdjustmentPeriod" tf:"budget_adjustment_period,omitempty"`
}

type NotificationInitParameters struct {

	// Comparison operator to use to evaluate the condition. Can be LESS_THAN, EQUAL_TO or GREATER_THAN.
	ComparisonOperator *string `json:"comparisonOperator,omitempty" tf:"comparison_operator,omitempty"`

	// What kind of budget value to notify on. Can be ACTUAL or FORECASTED
	NotificationType *string `json:"notificationType,omitempty" tf:"notification_type,omitempty"`

	// E-Mail addresses to notify. Either this or subscriber_sns_topic_arns is required.
	// +listType=set
	SubscriberEmailAddresses []*string `json:"subscriberEmailAddresses,omitempty" tf:"subscriber_email_addresses,omitempty"`

	// SNS topics to notify. Either this or subscriber_email_addresses is required.
	// +listType=set
	SubscriberSnsTopicArns []*string `json:"subscriberSnsTopicArns,omitempty" tf:"subscriber_sns_topic_arns,omitempty"`

	// Threshold when the notification should be sent.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// What kind of threshold is defined. Can be PERCENTAGE OR ABSOLUTE_VALUE.
	ThresholdType *string `json:"thresholdType,omitempty" tf:"threshold_type,omitempty"`
}

type NotificationObservation struct {

	// Comparison operator to use to evaluate the condition. Can be LESS_THAN, EQUAL_TO or GREATER_THAN.
	ComparisonOperator *string `json:"comparisonOperator,omitempty" tf:"comparison_operator,omitempty"`

	// What kind of budget value to notify on. Can be ACTUAL or FORECASTED
	NotificationType *string `json:"notificationType,omitempty" tf:"notification_type,omitempty"`

	// E-Mail addresses to notify. Either this or subscriber_sns_topic_arns is required.
	// +listType=set
	SubscriberEmailAddresses []*string `json:"subscriberEmailAddresses,omitempty" tf:"subscriber_email_addresses,omitempty"`

	// SNS topics to notify. Either this or subscriber_email_addresses is required.
	// +listType=set
	SubscriberSnsTopicArns []*string `json:"subscriberSnsTopicArns,omitempty" tf:"subscriber_sns_topic_arns,omitempty"`

	// Threshold when the notification should be sent.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// What kind of threshold is defined. Can be PERCENTAGE OR ABSOLUTE_VALUE.
	ThresholdType *string `json:"thresholdType,omitempty" tf:"threshold_type,omitempty"`
}

type NotificationParameters struct {

	// Comparison operator to use to evaluate the condition. Can be LESS_THAN, EQUAL_TO or GREATER_THAN.
	// +kubebuilder:validation:Optional
	ComparisonOperator *string `json:"comparisonOperator" tf:"comparison_operator,omitempty"`

	// What kind of budget value to notify on. Can be ACTUAL or FORECASTED
	// +kubebuilder:validation:Optional
	NotificationType *string `json:"notificationType" tf:"notification_type,omitempty"`

	// E-Mail addresses to notify. Either this or subscriber_sns_topic_arns is required.
	// +kubebuilder:validation:Optional
	// +listType=set
	SubscriberEmailAddresses []*string `json:"subscriberEmailAddresses,omitempty" tf:"subscriber_email_addresses,omitempty"`

	// SNS topics to notify. Either this or subscriber_email_addresses is required.
	// +kubebuilder:validation:Optional
	// +listType=set
	SubscriberSnsTopicArns []*string `json:"subscriberSnsTopicArns,omitempty" tf:"subscriber_sns_topic_arns,omitempty"`

	// Threshold when the notification should be sent.
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// What kind of threshold is defined. Can be PERCENTAGE OR ABSOLUTE_VALUE.
	// +kubebuilder:validation:Optional
	ThresholdType *string `json:"thresholdType" tf:"threshold_type,omitempty"`
}

type PlannedLimitInitParameters struct {

	// The amount of cost or usage being measured for a budget.
	Amount *string `json:"amount,omitempty" tf:"amount,omitempty"`

	// The start time of the budget limit. Format: 2017-01-01_12:00. See PlannedBudgetLimits documentation.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See Spend documentation.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type PlannedLimitObservation struct {

	// The amount of cost or usage being measured for a budget.
	Amount *string `json:"amount,omitempty" tf:"amount,omitempty"`

	// The start time of the budget limit. Format: 2017-01-01_12:00. See PlannedBudgetLimits documentation.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See Spend documentation.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type PlannedLimitParameters struct {

	// The amount of cost or usage being measured for a budget.
	// +kubebuilder:validation:Optional
	Amount *string `json:"amount" tf:"amount,omitempty"`

	// The start time of the budget limit. Format: 2017-01-01_12:00. See PlannedBudgetLimits documentation.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`

	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See Spend documentation.
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit" tf:"unit,omitempty"`
}

// BudgetSpec defines the desired state of Budget
type BudgetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BudgetParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider BudgetInitParameters `json:"initProvider,omitempty"`
}

// BudgetStatus defines the observed state of Budget.
type BudgetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BudgetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Budget is the Schema for the Budgets API. Provides a budgets budget resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Budget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.budgetType) || (has(self.initProvider) && has(self.initProvider.budgetType))",message="spec.forProvider.budgetType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.timeUnit) || (has(self.initProvider) && has(self.initProvider.timeUnit))",message="spec.forProvider.timeUnit is a required parameter"
	Spec   BudgetSpec   `json:"spec"`
	Status BudgetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BudgetList contains a list of Budgets
type BudgetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Budget `json:"items"`
}

// Repository type metadata.
var (
	Budget_Kind             = "Budget"
	Budget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Budget_Kind}.String()
	Budget_KindAPIVersion   = Budget_Kind + "." + CRDGroupVersion.String()
	Budget_GroupVersionKind = CRDGroupVersion.WithKind(Budget_Kind)
)

func init() {
	SchemeBuilder.Register(&Budget{}, &BudgetList{})
}
