# TournamenthttpCreateContestPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedRegions** | Pointer to [**[]ContestsRegionRequest**](ContestsRegionRequest.md) | AllowedRegions is available only for a worldwide Space. Entries narrow eligibility alongside the preset; empty adds no restriction. | [optional] 
**Answers** | Pointer to [**[]GuessGuessAnswer**](GuessGuessAnswer.md) |  | [optional] 
**AttemptLimit** | Pointer to **int32** |  | [optional] 
**BoardId** | Pointer to **string** |  | [optional] 
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

### NewTournamenthttpCreateContestPayload

`func NewTournamenthttpCreateContestPayload(reward DomainReward, sourceType string, spaceId string, timing DomainTiming, title string, ) *TournamenthttpCreateContestPayload`

NewTournamenthttpCreateContestPayload instantiates a new TournamenthttpCreateContestPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTournamenthttpCreateContestPayloadWithDefaults

`func NewTournamenthttpCreateContestPayloadWithDefaults() *TournamenthttpCreateContestPayload`

NewTournamenthttpCreateContestPayloadWithDefaults instantiates a new TournamenthttpCreateContestPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedRegions

`func (o *TournamenthttpCreateContestPayload) GetAllowedRegions() []ContestsRegionRequest`

GetAllowedRegions returns the AllowedRegions field if non-nil, zero value otherwise.

### GetAllowedRegionsOk

`func (o *TournamenthttpCreateContestPayload) GetAllowedRegionsOk() (*[]ContestsRegionRequest, bool)`

GetAllowedRegionsOk returns a tuple with the AllowedRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRegions

`func (o *TournamenthttpCreateContestPayload) SetAllowedRegions(v []ContestsRegionRequest)`

SetAllowedRegions sets AllowedRegions field to given value.

### HasAllowedRegions

`func (o *TournamenthttpCreateContestPayload) HasAllowedRegions() bool`

HasAllowedRegions returns a boolean if a field has been set.

### GetAnswers

`func (o *TournamenthttpCreateContestPayload) GetAnswers() []GuessGuessAnswer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *TournamenthttpCreateContestPayload) GetAnswersOk() (*[]GuessGuessAnswer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *TournamenthttpCreateContestPayload) SetAnswers(v []GuessGuessAnswer)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *TournamenthttpCreateContestPayload) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetAttemptLimit

`func (o *TournamenthttpCreateContestPayload) GetAttemptLimit() int32`

GetAttemptLimit returns the AttemptLimit field if non-nil, zero value otherwise.

### GetAttemptLimitOk

`func (o *TournamenthttpCreateContestPayload) GetAttemptLimitOk() (*int32, bool)`

GetAttemptLimitOk returns a tuple with the AttemptLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptLimit

`func (o *TournamenthttpCreateContestPayload) SetAttemptLimit(v int32)`

SetAttemptLimit sets AttemptLimit field to given value.

### HasAttemptLimit

`func (o *TournamenthttpCreateContestPayload) HasAttemptLimit() bool`

HasAttemptLimit returns a boolean if a field has been set.

### GetBoardId

`func (o *TournamenthttpCreateContestPayload) GetBoardId() string`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *TournamenthttpCreateContestPayload) GetBoardIdOk() (*string, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *TournamenthttpCreateContestPayload) SetBoardId(v string)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *TournamenthttpCreateContestPayload) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetCategory

`func (o *TournamenthttpCreateContestPayload) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TournamenthttpCreateContestPayload) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TournamenthttpCreateContestPayload) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TournamenthttpCreateContestPayload) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClaimScope

`func (o *TournamenthttpCreateContestPayload) GetClaimScope() string`

GetClaimScope returns the ClaimScope field if non-nil, zero value otherwise.

### GetClaimScopeOk

`func (o *TournamenthttpCreateContestPayload) GetClaimScopeOk() (*string, bool)`

GetClaimScopeOk returns a tuple with the ClaimScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimScope

`func (o *TournamenthttpCreateContestPayload) SetClaimScope(v string)`

SetClaimScope sets ClaimScope field to given value.

### HasClaimScope

`func (o *TournamenthttpCreateContestPayload) HasClaimScope() bool`

HasClaimScope returns a boolean if a field has been set.

### GetClueStrategy

`func (o *TournamenthttpCreateContestPayload) GetClueStrategy() string`

GetClueStrategy returns the ClueStrategy field if non-nil, zero value otherwise.

### GetClueStrategyOk

`func (o *TournamenthttpCreateContestPayload) GetClueStrategyOk() (*string, bool)`

GetClueStrategyOk returns a tuple with the ClueStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClueStrategy

`func (o *TournamenthttpCreateContestPayload) SetClueStrategy(v string)`

SetClueStrategy sets ClueStrategy field to given value.

### HasClueStrategy

`func (o *TournamenthttpCreateContestPayload) HasClueStrategy() bool`

HasClueStrategy returns a boolean if a field has been set.

### GetDescription

`func (o *TournamenthttpCreateContestPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TournamenthttpCreateContestPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TournamenthttpCreateContestPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TournamenthttpCreateContestPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDifficulty

`func (o *TournamenthttpCreateContestPayload) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *TournamenthttpCreateContestPayload) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *TournamenthttpCreateContestPayload) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *TournamenthttpCreateContestPayload) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetEnrollmentTiming

`func (o *TournamenthttpCreateContestPayload) GetEnrollmentTiming() DomainTiming`

GetEnrollmentTiming returns the EnrollmentTiming field if non-nil, zero value otherwise.

### GetEnrollmentTimingOk

`func (o *TournamenthttpCreateContestPayload) GetEnrollmentTimingOk() (*DomainTiming, bool)`

GetEnrollmentTimingOk returns a tuple with the EnrollmentTiming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentTiming

`func (o *TournamenthttpCreateContestPayload) SetEnrollmentTiming(v DomainTiming)`

SetEnrollmentTiming sets EnrollmentTiming field to given value.

### HasEnrollmentTiming

`func (o *TournamenthttpCreateContestPayload) HasEnrollmentTiming() bool`

HasEnrollmentTiming returns a boolean if a field has been set.

### GetGroupingType

`func (o *TournamenthttpCreateContestPayload) GetGroupingType() string`

GetGroupingType returns the GroupingType field if non-nil, zero value otherwise.

### GetGroupingTypeOk

`func (o *TournamenthttpCreateContestPayload) GetGroupingTypeOk() (*string, bool)`

GetGroupingTypeOk returns a tuple with the GroupingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingType

`func (o *TournamenthttpCreateContestPayload) SetGroupingType(v string)`

SetGroupingType sets GroupingType field to given value.

### HasGroupingType

`func (o *TournamenthttpCreateContestPayload) HasGroupingType() bool`

HasGroupingType returns a boolean if a field has been set.

### GetJudgingMethod

`func (o *TournamenthttpCreateContestPayload) GetJudgingMethod() string`

GetJudgingMethod returns the JudgingMethod field if non-nil, zero value otherwise.

### GetJudgingMethodOk

`func (o *TournamenthttpCreateContestPayload) GetJudgingMethodOk() (*string, bool)`

GetJudgingMethodOk returns a tuple with the JudgingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgingMethod

`func (o *TournamenthttpCreateContestPayload) SetJudgingMethod(v string)`

SetJudgingMethod sets JudgingMethod field to given value.

### HasJudgingMethod

`func (o *TournamenthttpCreateContestPayload) HasJudgingMethod() bool`

HasJudgingMethod returns a boolean if a field has been set.

### GetJurorSelectionMode

`func (o *TournamenthttpCreateContestPayload) GetJurorSelectionMode() string`

GetJurorSelectionMode returns the JurorSelectionMode field if non-nil, zero value otherwise.

### GetJurorSelectionModeOk

`func (o *TournamenthttpCreateContestPayload) GetJurorSelectionModeOk() (*string, bool)`

GetJurorSelectionModeOk returns a tuple with the JurorSelectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurorSelectionMode

`func (o *TournamenthttpCreateContestPayload) SetJurorSelectionMode(v string)`

SetJurorSelectionMode sets JurorSelectionMode field to given value.

### HasJurorSelectionMode

`func (o *TournamenthttpCreateContestPayload) HasJurorSelectionMode() bool`

HasJurorSelectionMode returns a boolean if a field has been set.

### GetMatchingRule

`func (o *TournamenthttpCreateContestPayload) GetMatchingRule() string`

GetMatchingRule returns the MatchingRule field if non-nil, zero value otherwise.

### GetMatchingRuleOk

`func (o *TournamenthttpCreateContestPayload) GetMatchingRuleOk() (*string, bool)`

GetMatchingRuleOk returns a tuple with the MatchingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingRule

`func (o *TournamenthttpCreateContestPayload) SetMatchingRule(v string)`

SetMatchingRule sets MatchingRule field to given value.

### HasMatchingRule

`func (o *TournamenthttpCreateContestPayload) HasMatchingRule() bool`

HasMatchingRule returns a boolean if a field has been set.

### GetMaxSelections

`func (o *TournamenthttpCreateContestPayload) GetMaxSelections() int32`

GetMaxSelections returns the MaxSelections field if non-nil, zero value otherwise.

### GetMaxSelectionsOk

`func (o *TournamenthttpCreateContestPayload) GetMaxSelectionsOk() (*int32, bool)`

GetMaxSelectionsOk returns a tuple with the MaxSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSelections

`func (o *TournamenthttpCreateContestPayload) SetMaxSelections(v int32)`

SetMaxSelections sets MaxSelections field to given value.

### HasMaxSelections

`func (o *TournamenthttpCreateContestPayload) HasMaxSelections() bool`

HasMaxSelections returns a boolean if a field has been set.

### GetNoOfGroups

`func (o *TournamenthttpCreateContestPayload) GetNoOfGroups() int32`

GetNoOfGroups returns the NoOfGroups field if non-nil, zero value otherwise.

### GetNoOfGroupsOk

`func (o *TournamenthttpCreateContestPayload) GetNoOfGroupsOk() (*int32, bool)`

GetNoOfGroupsOk returns a tuple with the NoOfGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfGroups

`func (o *TournamenthttpCreateContestPayload) SetNoOfGroups(v int32)`

SetNoOfGroups sets NoOfGroups field to given value.

### HasNoOfGroups

`func (o *TournamenthttpCreateContestPayload) HasNoOfGroups() bool`

HasNoOfGroups returns a boolean if a field has been set.

### GetNoOfPuzzles

`func (o *TournamenthttpCreateContestPayload) GetNoOfPuzzles() int32`

GetNoOfPuzzles returns the NoOfPuzzles field if non-nil, zero value otherwise.

### GetNoOfPuzzlesOk

`func (o *TournamenthttpCreateContestPayload) GetNoOfPuzzlesOk() (*int32, bool)`

GetNoOfPuzzlesOk returns a tuple with the NoOfPuzzles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfPuzzles

`func (o *TournamenthttpCreateContestPayload) SetNoOfPuzzles(v int32)`

SetNoOfPuzzles sets NoOfPuzzles field to given value.

### HasNoOfPuzzles

`func (o *TournamenthttpCreateContestPayload) HasNoOfPuzzles() bool`

HasNoOfPuzzles returns a boolean if a field has been set.

### GetOptions

`func (o *TournamenthttpCreateContestPayload) GetOptions() []VoteVoteOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *TournamenthttpCreateContestPayload) GetOptionsOk() (*[]VoteVoteOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *TournamenthttpCreateContestPayload) SetOptions(v []VoteVoteOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *TournamenthttpCreateContestPayload) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPrompt

`func (o *TournamenthttpCreateContestPayload) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *TournamenthttpCreateContestPayload) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *TournamenthttpCreateContestPayload) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *TournamenthttpCreateContestPayload) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetQuorum

`func (o *TournamenthttpCreateContestPayload) GetQuorum() int32`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *TournamenthttpCreateContestPayload) GetQuorumOk() (*int32, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *TournamenthttpCreateContestPayload) SetQuorum(v int32)`

SetQuorum sets Quorum field to given value.

### HasQuorum

`func (o *TournamenthttpCreateContestPayload) HasQuorum() bool`

HasQuorum returns a boolean if a field has been set.

### GetRatingModel

`func (o *TournamenthttpCreateContestPayload) GetRatingModel() string`

GetRatingModel returns the RatingModel field if non-nil, zero value otherwise.

### GetRatingModelOk

`func (o *TournamenthttpCreateContestPayload) GetRatingModelOk() (*string, bool)`

GetRatingModelOk returns a tuple with the RatingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingModel

`func (o *TournamenthttpCreateContestPayload) SetRatingModel(v string)`

SetRatingModel sets RatingModel field to given value.

### HasRatingModel

`func (o *TournamenthttpCreateContestPayload) HasRatingModel() bool`

HasRatingModel returns a boolean if a field has been set.

### GetReward

`func (o *TournamenthttpCreateContestPayload) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *TournamenthttpCreateContestPayload) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *TournamenthttpCreateContestPayload) SetReward(v DomainReward)`

SetReward sets Reward field to given value.


### GetRoomSize

`func (o *TournamenthttpCreateContestPayload) GetRoomSize() int32`

GetRoomSize returns the RoomSize field if non-nil, zero value otherwise.

### GetRoomSizeOk

`func (o *TournamenthttpCreateContestPayload) GetRoomSizeOk() (*int32, bool)`

GetRoomSizeOk returns a tuple with the RoomSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomSize

`func (o *TournamenthttpCreateContestPayload) SetRoomSize(v int32)`

SetRoomSize sets RoomSize field to given value.

### HasRoomSize

`func (o *TournamenthttpCreateContestPayload) HasRoomSize() bool`

HasRoomSize returns a boolean if a field has been set.

### GetRoomVisibility

`func (o *TournamenthttpCreateContestPayload) GetRoomVisibility() string`

GetRoomVisibility returns the RoomVisibility field if non-nil, zero value otherwise.

### GetRoomVisibilityOk

`func (o *TournamenthttpCreateContestPayload) GetRoomVisibilityOk() (*string, bool)`

GetRoomVisibilityOk returns a tuple with the RoomVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomVisibility

`func (o *TournamenthttpCreateContestPayload) SetRoomVisibility(v string)`

SetRoomVisibility sets RoomVisibility field to given value.

### HasRoomVisibility

`func (o *TournamenthttpCreateContestPayload) HasRoomVisibility() bool`

HasRoomVisibility returns a boolean if a field has been set.

### GetRulePreset

`func (o *TournamenthttpCreateContestPayload) GetRulePreset() string`

GetRulePreset returns the RulePreset field if non-nil, zero value otherwise.

### GetRulePresetOk

`func (o *TournamenthttpCreateContestPayload) GetRulePresetOk() (*string, bool)`

GetRulePresetOk returns a tuple with the RulePreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulePreset

`func (o *TournamenthttpCreateContestPayload) SetRulePreset(v string)`

SetRulePreset sets RulePreset field to given value.

### HasRulePreset

`func (o *TournamenthttpCreateContestPayload) HasRulePreset() bool`

HasRulePreset returns a boolean if a field has been set.

### GetSourceId

`func (o *TournamenthttpCreateContestPayload) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *TournamenthttpCreateContestPayload) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *TournamenthttpCreateContestPayload) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *TournamenthttpCreateContestPayload) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceType

`func (o *TournamenthttpCreateContestPayload) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TournamenthttpCreateContestPayload) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TournamenthttpCreateContestPayload) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetSourceVariant

`func (o *TournamenthttpCreateContestPayload) GetSourceVariant() string`

GetSourceVariant returns the SourceVariant field if non-nil, zero value otherwise.

### GetSourceVariantOk

`func (o *TournamenthttpCreateContestPayload) GetSourceVariantOk() (*string, bool)`

GetSourceVariantOk returns a tuple with the SourceVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVariant

`func (o *TournamenthttpCreateContestPayload) SetSourceVariant(v string)`

SetSourceVariant sets SourceVariant field to given value.

### HasSourceVariant

`func (o *TournamenthttpCreateContestPayload) HasSourceVariant() bool`

HasSourceVariant returns a boolean if a field has been set.

### GetSpaceId

`func (o *TournamenthttpCreateContestPayload) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *TournamenthttpCreateContestPayload) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *TournamenthttpCreateContestPayload) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetSponsorshipPackageId

`func (o *TournamenthttpCreateContestPayload) GetSponsorshipPackageId() string`

GetSponsorshipPackageId returns the SponsorshipPackageId field if non-nil, zero value otherwise.

### GetSponsorshipPackageIdOk

`func (o *TournamenthttpCreateContestPayload) GetSponsorshipPackageIdOk() (*string, bool)`

GetSponsorshipPackageIdOk returns a tuple with the SponsorshipPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorshipPackageId

`func (o *TournamenthttpCreateContestPayload) SetSponsorshipPackageId(v string)`

SetSponsorshipPackageId sets SponsorshipPackageId field to given value.

### HasSponsorshipPackageId

`func (o *TournamenthttpCreateContestPayload) HasSponsorshipPackageId() bool`

HasSponsorshipPackageId returns a boolean if a field has been set.

### GetSubmissionPolicy

`func (o *TournamenthttpCreateContestPayload) GetSubmissionPolicy() string`

GetSubmissionPolicy returns the SubmissionPolicy field if non-nil, zero value otherwise.

### GetSubmissionPolicyOk

`func (o *TournamenthttpCreateContestPayload) GetSubmissionPolicyOk() (*string, bool)`

GetSubmissionPolicyOk returns a tuple with the SubmissionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionPolicy

`func (o *TournamenthttpCreateContestPayload) SetSubmissionPolicy(v string)`

SetSubmissionPolicy sets SubmissionPolicy field to given value.

### HasSubmissionPolicy

`func (o *TournamenthttpCreateContestPayload) HasSubmissionPolicy() bool`

HasSubmissionPolicy returns a boolean if a field has been set.

### GetTeamSetId

`func (o *TournamenthttpCreateContestPayload) GetTeamSetId() string`

GetTeamSetId returns the TeamSetId field if non-nil, zero value otherwise.

### GetTeamSetIdOk

`func (o *TournamenthttpCreateContestPayload) GetTeamSetIdOk() (*string, bool)`

GetTeamSetIdOk returns a tuple with the TeamSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamSetId

`func (o *TournamenthttpCreateContestPayload) SetTeamSetId(v string)`

SetTeamSetId sets TeamSetId field to given value.

### HasTeamSetId

`func (o *TournamenthttpCreateContestPayload) HasTeamSetId() bool`

HasTeamSetId returns a boolean if a field has been set.

### GetTieBreak

`func (o *TournamenthttpCreateContestPayload) GetTieBreak() string`

GetTieBreak returns the TieBreak field if non-nil, zero value otherwise.

### GetTieBreakOk

`func (o *TournamenthttpCreateContestPayload) GetTieBreakOk() (*string, bool)`

GetTieBreakOk returns a tuple with the TieBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieBreak

`func (o *TournamenthttpCreateContestPayload) SetTieBreak(v string)`

SetTieBreak sets TieBreak field to given value.

### HasTieBreak

`func (o *TournamenthttpCreateContestPayload) HasTieBreak() bool`

HasTieBreak returns a boolean if a field has been set.

### GetTiming

`func (o *TournamenthttpCreateContestPayload) GetTiming() DomainTiming`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *TournamenthttpCreateContestPayload) GetTimingOk() (*DomainTiming, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *TournamenthttpCreateContestPayload) SetTiming(v DomainTiming)`

SetTiming sets Timing field to given value.


### GetTitle

`func (o *TournamenthttpCreateContestPayload) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TournamenthttpCreateContestPayload) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TournamenthttpCreateContestPayload) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVoteVisibility

`func (o *TournamenthttpCreateContestPayload) GetVoteVisibility() string`

GetVoteVisibility returns the VoteVisibility field if non-nil, zero value otherwise.

### GetVoteVisibilityOk

`func (o *TournamenthttpCreateContestPayload) GetVoteVisibilityOk() (*string, bool)`

GetVoteVisibilityOk returns a tuple with the VoteVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteVisibility

`func (o *TournamenthttpCreateContestPayload) SetVoteVisibility(v string)`

SetVoteVisibility sets VoteVisibility field to given value.

### HasVoteVisibility

`func (o *TournamenthttpCreateContestPayload) HasVoteVisibility() bool`

HasVoteVisibility returns a boolean if a field has been set.

### GetWinnerCount

`func (o *TournamenthttpCreateContestPayload) GetWinnerCount() int32`

GetWinnerCount returns the WinnerCount field if non-nil, zero value otherwise.

### GetWinnerCountOk

`func (o *TournamenthttpCreateContestPayload) GetWinnerCountOk() (*int32, bool)`

GetWinnerCountOk returns a tuple with the WinnerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerCount

`func (o *TournamenthttpCreateContestPayload) SetWinnerCount(v int32)`

SetWinnerCount sets WinnerCount field to given value.

### HasWinnerCount

`func (o *TournamenthttpCreateContestPayload) HasWinnerCount() bool`

HasWinnerCount returns a boolean if a field has been set.

### GetWinnersPerGroup

`func (o *TournamenthttpCreateContestPayload) GetWinnersPerGroup() int32`

GetWinnersPerGroup returns the WinnersPerGroup field if non-nil, zero value otherwise.

### GetWinnersPerGroupOk

`func (o *TournamenthttpCreateContestPayload) GetWinnersPerGroupOk() (*int32, bool)`

GetWinnersPerGroupOk returns a tuple with the WinnersPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnersPerGroup

`func (o *TournamenthttpCreateContestPayload) SetWinnersPerGroup(v int32)`

SetWinnersPerGroup sets WinnersPerGroup field to given value.

### HasWinnersPerGroup

`func (o *TournamenthttpCreateContestPayload) HasWinnersPerGroup() bool`

HasWinnersPerGroup returns a boolean if a field has been set.

### GetWinningStrategy

`func (o *TournamenthttpCreateContestPayload) GetWinningStrategy() string`

GetWinningStrategy returns the WinningStrategy field if non-nil, zero value otherwise.

### GetWinningStrategyOk

`func (o *TournamenthttpCreateContestPayload) GetWinningStrategyOk() (*string, bool)`

GetWinningStrategyOk returns a tuple with the WinningStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinningStrategy

`func (o *TournamenthttpCreateContestPayload) SetWinningStrategy(v string)`

SetWinningStrategy sets WinningStrategy field to given value.

### HasWinningStrategy

`func (o *TournamenthttpCreateContestPayload) HasWinningStrategy() bool`

HasWinningStrategy returns a boolean if a field has been set.

### GetZoneAdvantage

`func (o *TournamenthttpCreateContestPayload) GetZoneAdvantage() int32`

GetZoneAdvantage returns the ZoneAdvantage field if non-nil, zero value otherwise.

### GetZoneAdvantageOk

`func (o *TournamenthttpCreateContestPayload) GetZoneAdvantageOk() (*int32, bool)`

GetZoneAdvantageOk returns a tuple with the ZoneAdvantage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAdvantage

`func (o *TournamenthttpCreateContestPayload) SetZoneAdvantage(v int32)`

SetZoneAdvantage sets ZoneAdvantage field to given value.

### HasZoneAdvantage

`func (o *TournamenthttpCreateContestPayload) HasZoneAdvantage() bool`

HasZoneAdvantage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


