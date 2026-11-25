# ContesthttpCreateContestPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedRegions** | Pointer to [**[]ContestsRegionRequest**](ContestsRegionRequest.md) | AllowedRegions is available only for a worldwide Space. Entries narrow eligibility alongside the preset; empty adds no restriction. | [optional] 
**Answers** | Pointer to [**[]GuessGuessAnswer**](GuessGuessAnswer.md) |  | [optional] 
**AttemptLimit** | Pointer to **int32** |  | [optional] 
**BoardId** | Pointer to **string** |  | [optional] 
**BonusRewards** | Pointer to [**[]DomainBonusReward**](DomainBonusReward.md) |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ClaimScope** | Pointer to **string** |  | [optional] 
**ClueStrategy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Difficulty** | Pointer to **string** |  | [optional] 
**EnrollmentTiming** | Pointer to [**DomainTiming**](DomainTiming.md) | EnrollmentTiming opts this contest into pre-enrollment (ADR 0058): any contest type may carry one, regardless of whether its type structurally requires it. When omitted on a type that requires enrollment, the platform defaults one (AsNewContest); nil otherwise. | [optional] 
**GroupingType** | Pointer to **string** |  | [optional] 
**JudgingMethod** | Pointer to **string** |  | [optional] 
**JurorSelectionMode** | Pointer to **string** |  | [optional] 
**MatchingRule** | Pointer to **string** |  | [optional] 
**MaxSelections** | Pointer to **int32** |  | [optional] 
**NoOfGroups** | Pointer to **int32** |  | [optional] 
**NoOfPuzzles** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**[]VoteVoteOption**](VoteVoteOption.md) |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**Quorum** | Pointer to **int32** |  | [optional] 
**RatingModel** | Pointer to **string** |  | [optional] 
**Reward** | [**DomainReward**](DomainReward.md) |  | 
**RoomSize** | Pointer to **int32** |  | [optional] 
**RoomVisibility** | Pointer to **string** |  | [optional] 
**RulePreset** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**SourceType** | **string** |  | 
**SourceVariant** | Pointer to **string** |  | [optional] 
**SpaceId** | **string** |  | 
**SponsorshipPackageId** | Pointer to **string** |  | [optional] 
**SubmissionPolicy** | Pointer to **string** |  | [optional] 
**TeamSetId** | Pointer to **string** |  | [optional] 
**TieBreak** | Pointer to **string** |  | [optional] 
**Timing** | [**DomainTiming**](DomainTiming.md) |  | 
**Title** | **string** |  | 
**VoteVisibility** | Pointer to **string** |  | [optional] 
**WinnerCount** | Pointer to **int32** |  | [optional] 
**WinnersPerGroup** | Pointer to **int32** |  | [optional] 
**WinningStrategy** | Pointer to **string** |  | [optional] 
**ZoneAdvantage** | Pointer to **int32** | ZoneAdvantage configures group-scoped score deductions for MCQ/Moment only. | [optional] 

## Methods

### NewContesthttpCreateContestPayload

`func NewContesthttpCreateContestPayload(reward DomainReward, sourceType string, spaceId string, timing DomainTiming, title string, ) *ContesthttpCreateContestPayload`

NewContesthttpCreateContestPayload instantiates a new ContesthttpCreateContestPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContesthttpCreateContestPayloadWithDefaults

`func NewContesthttpCreateContestPayloadWithDefaults() *ContesthttpCreateContestPayload`

NewContesthttpCreateContestPayloadWithDefaults instantiates a new ContesthttpCreateContestPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedRegions

`func (o *ContesthttpCreateContestPayload) GetAllowedRegions() []ContestsRegionRequest`

GetAllowedRegions returns the AllowedRegions field if non-nil, zero value otherwise.

### GetAllowedRegionsOk

`func (o *ContesthttpCreateContestPayload) GetAllowedRegionsOk() (*[]ContestsRegionRequest, bool)`

GetAllowedRegionsOk returns a tuple with the AllowedRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRegions

`func (o *ContesthttpCreateContestPayload) SetAllowedRegions(v []ContestsRegionRequest)`

SetAllowedRegions sets AllowedRegions field to given value.

### HasAllowedRegions

`func (o *ContesthttpCreateContestPayload) HasAllowedRegions() bool`

HasAllowedRegions returns a boolean if a field has been set.

### GetAnswers

`func (o *ContesthttpCreateContestPayload) GetAnswers() []GuessGuessAnswer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *ContesthttpCreateContestPayload) GetAnswersOk() (*[]GuessGuessAnswer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *ContesthttpCreateContestPayload) SetAnswers(v []GuessGuessAnswer)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *ContesthttpCreateContestPayload) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetAttemptLimit

`func (o *ContesthttpCreateContestPayload) GetAttemptLimit() int32`

GetAttemptLimit returns the AttemptLimit field if non-nil, zero value otherwise.

### GetAttemptLimitOk

`func (o *ContesthttpCreateContestPayload) GetAttemptLimitOk() (*int32, bool)`

GetAttemptLimitOk returns a tuple with the AttemptLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptLimit

`func (o *ContesthttpCreateContestPayload) SetAttemptLimit(v int32)`

SetAttemptLimit sets AttemptLimit field to given value.

### HasAttemptLimit

`func (o *ContesthttpCreateContestPayload) HasAttemptLimit() bool`

HasAttemptLimit returns a boolean if a field has been set.

### GetBoardId

`func (o *ContesthttpCreateContestPayload) GetBoardId() string`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *ContesthttpCreateContestPayload) GetBoardIdOk() (*string, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *ContesthttpCreateContestPayload) SetBoardId(v string)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *ContesthttpCreateContestPayload) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetBonusRewards

`func (o *ContesthttpCreateContestPayload) GetBonusRewards() []DomainBonusReward`

GetBonusRewards returns the BonusRewards field if non-nil, zero value otherwise.

### GetBonusRewardsOk

`func (o *ContesthttpCreateContestPayload) GetBonusRewardsOk() (*[]DomainBonusReward, bool)`

GetBonusRewardsOk returns a tuple with the BonusRewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonusRewards

`func (o *ContesthttpCreateContestPayload) SetBonusRewards(v []DomainBonusReward)`

SetBonusRewards sets BonusRewards field to given value.

### HasBonusRewards

`func (o *ContesthttpCreateContestPayload) HasBonusRewards() bool`

HasBonusRewards returns a boolean if a field has been set.

### GetCategory

`func (o *ContesthttpCreateContestPayload) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ContesthttpCreateContestPayload) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ContesthttpCreateContestPayload) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ContesthttpCreateContestPayload) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClaimScope

`func (o *ContesthttpCreateContestPayload) GetClaimScope() string`

GetClaimScope returns the ClaimScope field if non-nil, zero value otherwise.

### GetClaimScopeOk

`func (o *ContesthttpCreateContestPayload) GetClaimScopeOk() (*string, bool)`

GetClaimScopeOk returns a tuple with the ClaimScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimScope

`func (o *ContesthttpCreateContestPayload) SetClaimScope(v string)`

SetClaimScope sets ClaimScope field to given value.

### HasClaimScope

`func (o *ContesthttpCreateContestPayload) HasClaimScope() bool`

HasClaimScope returns a boolean if a field has been set.

### GetClueStrategy

`func (o *ContesthttpCreateContestPayload) GetClueStrategy() string`

GetClueStrategy returns the ClueStrategy field if non-nil, zero value otherwise.

### GetClueStrategyOk

`func (o *ContesthttpCreateContestPayload) GetClueStrategyOk() (*string, bool)`

GetClueStrategyOk returns a tuple with the ClueStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClueStrategy

`func (o *ContesthttpCreateContestPayload) SetClueStrategy(v string)`

SetClueStrategy sets ClueStrategy field to given value.

### HasClueStrategy

`func (o *ContesthttpCreateContestPayload) HasClueStrategy() bool`

HasClueStrategy returns a boolean if a field has been set.

### GetDescription

`func (o *ContesthttpCreateContestPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContesthttpCreateContestPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContesthttpCreateContestPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContesthttpCreateContestPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDifficulty

`func (o *ContesthttpCreateContestPayload) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *ContesthttpCreateContestPayload) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *ContesthttpCreateContestPayload) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *ContesthttpCreateContestPayload) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetEnrollmentTiming

`func (o *ContesthttpCreateContestPayload) GetEnrollmentTiming() DomainTiming`

GetEnrollmentTiming returns the EnrollmentTiming field if non-nil, zero value otherwise.

### GetEnrollmentTimingOk

`func (o *ContesthttpCreateContestPayload) GetEnrollmentTimingOk() (*DomainTiming, bool)`

GetEnrollmentTimingOk returns a tuple with the EnrollmentTiming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentTiming

`func (o *ContesthttpCreateContestPayload) SetEnrollmentTiming(v DomainTiming)`

SetEnrollmentTiming sets EnrollmentTiming field to given value.

### HasEnrollmentTiming

`func (o *ContesthttpCreateContestPayload) HasEnrollmentTiming() bool`

HasEnrollmentTiming returns a boolean if a field has been set.

### GetGroupingType

`func (o *ContesthttpCreateContestPayload) GetGroupingType() string`

GetGroupingType returns the GroupingType field if non-nil, zero value otherwise.

### GetGroupingTypeOk

`func (o *ContesthttpCreateContestPayload) GetGroupingTypeOk() (*string, bool)`

GetGroupingTypeOk returns a tuple with the GroupingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingType

`func (o *ContesthttpCreateContestPayload) SetGroupingType(v string)`

SetGroupingType sets GroupingType field to given value.

### HasGroupingType

`func (o *ContesthttpCreateContestPayload) HasGroupingType() bool`

HasGroupingType returns a boolean if a field has been set.

### GetJudgingMethod

`func (o *ContesthttpCreateContestPayload) GetJudgingMethod() string`

GetJudgingMethod returns the JudgingMethod field if non-nil, zero value otherwise.

### GetJudgingMethodOk

`func (o *ContesthttpCreateContestPayload) GetJudgingMethodOk() (*string, bool)`

GetJudgingMethodOk returns a tuple with the JudgingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgingMethod

`func (o *ContesthttpCreateContestPayload) SetJudgingMethod(v string)`

SetJudgingMethod sets JudgingMethod field to given value.

### HasJudgingMethod

`func (o *ContesthttpCreateContestPayload) HasJudgingMethod() bool`

HasJudgingMethod returns a boolean if a field has been set.

### GetJurorSelectionMode

`func (o *ContesthttpCreateContestPayload) GetJurorSelectionMode() string`

GetJurorSelectionMode returns the JurorSelectionMode field if non-nil, zero value otherwise.

### GetJurorSelectionModeOk

`func (o *ContesthttpCreateContestPayload) GetJurorSelectionModeOk() (*string, bool)`

GetJurorSelectionModeOk returns a tuple with the JurorSelectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurorSelectionMode

`func (o *ContesthttpCreateContestPayload) SetJurorSelectionMode(v string)`

SetJurorSelectionMode sets JurorSelectionMode field to given value.

### HasJurorSelectionMode

`func (o *ContesthttpCreateContestPayload) HasJurorSelectionMode() bool`

HasJurorSelectionMode returns a boolean if a field has been set.

### GetMatchingRule

`func (o *ContesthttpCreateContestPayload) GetMatchingRule() string`

GetMatchingRule returns the MatchingRule field if non-nil, zero value otherwise.

### GetMatchingRuleOk

`func (o *ContesthttpCreateContestPayload) GetMatchingRuleOk() (*string, bool)`

GetMatchingRuleOk returns a tuple with the MatchingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingRule

`func (o *ContesthttpCreateContestPayload) SetMatchingRule(v string)`

SetMatchingRule sets MatchingRule field to given value.

### HasMatchingRule

`func (o *ContesthttpCreateContestPayload) HasMatchingRule() bool`

HasMatchingRule returns a boolean if a field has been set.

### GetMaxSelections

`func (o *ContesthttpCreateContestPayload) GetMaxSelections() int32`

GetMaxSelections returns the MaxSelections field if non-nil, zero value otherwise.

### GetMaxSelectionsOk

`func (o *ContesthttpCreateContestPayload) GetMaxSelectionsOk() (*int32, bool)`

GetMaxSelectionsOk returns a tuple with the MaxSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSelections

`func (o *ContesthttpCreateContestPayload) SetMaxSelections(v int32)`

SetMaxSelections sets MaxSelections field to given value.

### HasMaxSelections

`func (o *ContesthttpCreateContestPayload) HasMaxSelections() bool`

HasMaxSelections returns a boolean if a field has been set.

### GetNoOfGroups

`func (o *ContesthttpCreateContestPayload) GetNoOfGroups() int32`

GetNoOfGroups returns the NoOfGroups field if non-nil, zero value otherwise.

### GetNoOfGroupsOk

`func (o *ContesthttpCreateContestPayload) GetNoOfGroupsOk() (*int32, bool)`

GetNoOfGroupsOk returns a tuple with the NoOfGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfGroups

`func (o *ContesthttpCreateContestPayload) SetNoOfGroups(v int32)`

SetNoOfGroups sets NoOfGroups field to given value.

### HasNoOfGroups

`func (o *ContesthttpCreateContestPayload) HasNoOfGroups() bool`

HasNoOfGroups returns a boolean if a field has been set.

### GetNoOfPuzzles

`func (o *ContesthttpCreateContestPayload) GetNoOfPuzzles() int32`

GetNoOfPuzzles returns the NoOfPuzzles field if non-nil, zero value otherwise.

### GetNoOfPuzzlesOk

`func (o *ContesthttpCreateContestPayload) GetNoOfPuzzlesOk() (*int32, bool)`

GetNoOfPuzzlesOk returns a tuple with the NoOfPuzzles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfPuzzles

`func (o *ContesthttpCreateContestPayload) SetNoOfPuzzles(v int32)`

SetNoOfPuzzles sets NoOfPuzzles field to given value.

### HasNoOfPuzzles

`func (o *ContesthttpCreateContestPayload) HasNoOfPuzzles() bool`

HasNoOfPuzzles returns a boolean if a field has been set.

### GetOptions

`func (o *ContesthttpCreateContestPayload) GetOptions() []VoteVoteOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ContesthttpCreateContestPayload) GetOptionsOk() (*[]VoteVoteOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ContesthttpCreateContestPayload) SetOptions(v []VoteVoteOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ContesthttpCreateContestPayload) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPrompt

`func (o *ContesthttpCreateContestPayload) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *ContesthttpCreateContestPayload) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *ContesthttpCreateContestPayload) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *ContesthttpCreateContestPayload) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetQuorum

`func (o *ContesthttpCreateContestPayload) GetQuorum() int32`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *ContesthttpCreateContestPayload) GetQuorumOk() (*int32, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *ContesthttpCreateContestPayload) SetQuorum(v int32)`

SetQuorum sets Quorum field to given value.

### HasQuorum

`func (o *ContesthttpCreateContestPayload) HasQuorum() bool`

HasQuorum returns a boolean if a field has been set.

### GetRatingModel

`func (o *ContesthttpCreateContestPayload) GetRatingModel() string`

GetRatingModel returns the RatingModel field if non-nil, zero value otherwise.

### GetRatingModelOk

`func (o *ContesthttpCreateContestPayload) GetRatingModelOk() (*string, bool)`

GetRatingModelOk returns a tuple with the RatingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingModel

`func (o *ContesthttpCreateContestPayload) SetRatingModel(v string)`

SetRatingModel sets RatingModel field to given value.

### HasRatingModel

`func (o *ContesthttpCreateContestPayload) HasRatingModel() bool`

HasRatingModel returns a boolean if a field has been set.

### GetReward

`func (o *ContesthttpCreateContestPayload) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *ContesthttpCreateContestPayload) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *ContesthttpCreateContestPayload) SetReward(v DomainReward)`

SetReward sets Reward field to given value.


### GetRoomSize

`func (o *ContesthttpCreateContestPayload) GetRoomSize() int32`

GetRoomSize returns the RoomSize field if non-nil, zero value otherwise.

### GetRoomSizeOk

`func (o *ContesthttpCreateContestPayload) GetRoomSizeOk() (*int32, bool)`

GetRoomSizeOk returns a tuple with the RoomSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomSize

`func (o *ContesthttpCreateContestPayload) SetRoomSize(v int32)`

SetRoomSize sets RoomSize field to given value.

### HasRoomSize

`func (o *ContesthttpCreateContestPayload) HasRoomSize() bool`

HasRoomSize returns a boolean if a field has been set.

### GetRoomVisibility

`func (o *ContesthttpCreateContestPayload) GetRoomVisibility() string`

GetRoomVisibility returns the RoomVisibility field if non-nil, zero value otherwise.

### GetRoomVisibilityOk

`func (o *ContesthttpCreateContestPayload) GetRoomVisibilityOk() (*string, bool)`

GetRoomVisibilityOk returns a tuple with the RoomVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomVisibility

`func (o *ContesthttpCreateContestPayload) SetRoomVisibility(v string)`

SetRoomVisibility sets RoomVisibility field to given value.

### HasRoomVisibility

`func (o *ContesthttpCreateContestPayload) HasRoomVisibility() bool`

HasRoomVisibility returns a boolean if a field has been set.

### GetRulePreset

`func (o *ContesthttpCreateContestPayload) GetRulePreset() string`

GetRulePreset returns the RulePreset field if non-nil, zero value otherwise.

### GetRulePresetOk

`func (o *ContesthttpCreateContestPayload) GetRulePresetOk() (*string, bool)`

GetRulePresetOk returns a tuple with the RulePreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulePreset

`func (o *ContesthttpCreateContestPayload) SetRulePreset(v string)`

SetRulePreset sets RulePreset field to given value.

### HasRulePreset

`func (o *ContesthttpCreateContestPayload) HasRulePreset() bool`

HasRulePreset returns a boolean if a field has been set.

### GetSourceId

`func (o *ContesthttpCreateContestPayload) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ContesthttpCreateContestPayload) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ContesthttpCreateContestPayload) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ContesthttpCreateContestPayload) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceType

`func (o *ContesthttpCreateContestPayload) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ContesthttpCreateContestPayload) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ContesthttpCreateContestPayload) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetSourceVariant

`func (o *ContesthttpCreateContestPayload) GetSourceVariant() string`

GetSourceVariant returns the SourceVariant field if non-nil, zero value otherwise.

### GetSourceVariantOk

`func (o *ContesthttpCreateContestPayload) GetSourceVariantOk() (*string, bool)`

GetSourceVariantOk returns a tuple with the SourceVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVariant

`func (o *ContesthttpCreateContestPayload) SetSourceVariant(v string)`

SetSourceVariant sets SourceVariant field to given value.

### HasSourceVariant

`func (o *ContesthttpCreateContestPayload) HasSourceVariant() bool`

HasSourceVariant returns a boolean if a field has been set.

### GetSpaceId

`func (o *ContesthttpCreateContestPayload) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *ContesthttpCreateContestPayload) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *ContesthttpCreateContestPayload) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetSponsorshipPackageId

`func (o *ContesthttpCreateContestPayload) GetSponsorshipPackageId() string`

GetSponsorshipPackageId returns the SponsorshipPackageId field if non-nil, zero value otherwise.

### GetSponsorshipPackageIdOk

`func (o *ContesthttpCreateContestPayload) GetSponsorshipPackageIdOk() (*string, bool)`

GetSponsorshipPackageIdOk returns a tuple with the SponsorshipPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorshipPackageId

`func (o *ContesthttpCreateContestPayload) SetSponsorshipPackageId(v string)`

SetSponsorshipPackageId sets SponsorshipPackageId field to given value.

### HasSponsorshipPackageId

`func (o *ContesthttpCreateContestPayload) HasSponsorshipPackageId() bool`

HasSponsorshipPackageId returns a boolean if a field has been set.

### GetSubmissionPolicy

`func (o *ContesthttpCreateContestPayload) GetSubmissionPolicy() string`

GetSubmissionPolicy returns the SubmissionPolicy field if non-nil, zero value otherwise.

### GetSubmissionPolicyOk

`func (o *ContesthttpCreateContestPayload) GetSubmissionPolicyOk() (*string, bool)`

GetSubmissionPolicyOk returns a tuple with the SubmissionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionPolicy

`func (o *ContesthttpCreateContestPayload) SetSubmissionPolicy(v string)`

SetSubmissionPolicy sets SubmissionPolicy field to given value.

### HasSubmissionPolicy

`func (o *ContesthttpCreateContestPayload) HasSubmissionPolicy() bool`

HasSubmissionPolicy returns a boolean if a field has been set.

### GetTeamSetId

`func (o *ContesthttpCreateContestPayload) GetTeamSetId() string`

GetTeamSetId returns the TeamSetId field if non-nil, zero value otherwise.

### GetTeamSetIdOk

`func (o *ContesthttpCreateContestPayload) GetTeamSetIdOk() (*string, bool)`

GetTeamSetIdOk returns a tuple with the TeamSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamSetId

`func (o *ContesthttpCreateContestPayload) SetTeamSetId(v string)`

SetTeamSetId sets TeamSetId field to given value.

### HasTeamSetId

`func (o *ContesthttpCreateContestPayload) HasTeamSetId() bool`

HasTeamSetId returns a boolean if a field has been set.

### GetTieBreak

`func (o *ContesthttpCreateContestPayload) GetTieBreak() string`

GetTieBreak returns the TieBreak field if non-nil, zero value otherwise.

### GetTieBreakOk

`func (o *ContesthttpCreateContestPayload) GetTieBreakOk() (*string, bool)`

GetTieBreakOk returns a tuple with the TieBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieBreak

`func (o *ContesthttpCreateContestPayload) SetTieBreak(v string)`

SetTieBreak sets TieBreak field to given value.

### HasTieBreak

`func (o *ContesthttpCreateContestPayload) HasTieBreak() bool`

HasTieBreak returns a boolean if a field has been set.

### GetTiming

`func (o *ContesthttpCreateContestPayload) GetTiming() DomainTiming`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *ContesthttpCreateContestPayload) GetTimingOk() (*DomainTiming, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *ContesthttpCreateContestPayload) SetTiming(v DomainTiming)`

SetTiming sets Timing field to given value.


### GetTitle

`func (o *ContesthttpCreateContestPayload) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContesthttpCreateContestPayload) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContesthttpCreateContestPayload) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVoteVisibility

`func (o *ContesthttpCreateContestPayload) GetVoteVisibility() string`

GetVoteVisibility returns the VoteVisibility field if non-nil, zero value otherwise.

### GetVoteVisibilityOk

`func (o *ContesthttpCreateContestPayload) GetVoteVisibilityOk() (*string, bool)`

GetVoteVisibilityOk returns a tuple with the VoteVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteVisibility

`func (o *ContesthttpCreateContestPayload) SetVoteVisibility(v string)`

SetVoteVisibility sets VoteVisibility field to given value.

### HasVoteVisibility

`func (o *ContesthttpCreateContestPayload) HasVoteVisibility() bool`

HasVoteVisibility returns a boolean if a field has been set.

### GetWinnerCount

`func (o *ContesthttpCreateContestPayload) GetWinnerCount() int32`

GetWinnerCount returns the WinnerCount field if non-nil, zero value otherwise.

### GetWinnerCountOk

`func (o *ContesthttpCreateContestPayload) GetWinnerCountOk() (*int32, bool)`

GetWinnerCountOk returns a tuple with the WinnerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerCount

`func (o *ContesthttpCreateContestPayload) SetWinnerCount(v int32)`

SetWinnerCount sets WinnerCount field to given value.

### HasWinnerCount

`func (o *ContesthttpCreateContestPayload) HasWinnerCount() bool`

HasWinnerCount returns a boolean if a field has been set.

### GetWinnersPerGroup

`func (o *ContesthttpCreateContestPayload) GetWinnersPerGroup() int32`

GetWinnersPerGroup returns the WinnersPerGroup field if non-nil, zero value otherwise.

### GetWinnersPerGroupOk

`func (o *ContesthttpCreateContestPayload) GetWinnersPerGroupOk() (*int32, bool)`

GetWinnersPerGroupOk returns a tuple with the WinnersPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnersPerGroup

`func (o *ContesthttpCreateContestPayload) SetWinnersPerGroup(v int32)`

SetWinnersPerGroup sets WinnersPerGroup field to given value.

### HasWinnersPerGroup

`func (o *ContesthttpCreateContestPayload) HasWinnersPerGroup() bool`

HasWinnersPerGroup returns a boolean if a field has been set.

### GetWinningStrategy

`func (o *ContesthttpCreateContestPayload) GetWinningStrategy() string`

GetWinningStrategy returns the WinningStrategy field if non-nil, zero value otherwise.

### GetWinningStrategyOk

`func (o *ContesthttpCreateContestPayload) GetWinningStrategyOk() (*string, bool)`

GetWinningStrategyOk returns a tuple with the WinningStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinningStrategy

`func (o *ContesthttpCreateContestPayload) SetWinningStrategy(v string)`

SetWinningStrategy sets WinningStrategy field to given value.

### HasWinningStrategy

`func (o *ContesthttpCreateContestPayload) HasWinningStrategy() bool`

HasWinningStrategy returns a boolean if a field has been set.

### GetZoneAdvantage

`func (o *ContesthttpCreateContestPayload) GetZoneAdvantage() int32`

GetZoneAdvantage returns the ZoneAdvantage field if non-nil, zero value otherwise.

### GetZoneAdvantageOk

`func (o *ContesthttpCreateContestPayload) GetZoneAdvantageOk() (*int32, bool)`

GetZoneAdvantageOk returns a tuple with the ZoneAdvantage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAdvantage

`func (o *ContesthttpCreateContestPayload) SetZoneAdvantage(v int32)`

SetZoneAdvantage sets ZoneAdvantage field to given value.

### HasZoneAdvantage

`func (o *ContesthttpCreateContestPayload) HasZoneAdvantage() bool`

HasZoneAdvantage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


