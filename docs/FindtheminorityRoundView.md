# FindtheminorityRoundView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccusationSubmittedCount** | Pointer to **int32** |  | [optional] 
**AccusedAnswer** | Pointer to **string** |  | [optional] 
**AccusedIdentityId** | Pointer to **string** |  | [optional] 
**AnswerCounts** | Pointer to **map[string]int32** |  | [optional] 
**AnswerSubmittedCount** | Pointer to **int32** |  | [optional] 
**Answers** | Pointer to **map[string]string** |  | [optional] 
**AudienceBoard** | Pointer to [**VoteTallyView**](VoteTallyView.md) |  | [optional] 
**AudiencePollState** | Pointer to **string** |  | [optional] 
**AudienceVerdictRevealed** | Pointer to **bool** |  | [optional] 
**BoardToAudience** | Pointer to **bool** |  | [optional] 
**BoardToPlayers** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MinorityOption** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**[]VoteOption**](VoteOption.md) |  | [optional] 
**OutcomeKind** | Pointer to **string** |  | [optional] 
**OwnAccusation** | Pointer to **string** |  | [optional] 
**OwnAnswer** | Pointer to **string** |  | [optional] 
**OwnAudienceVote** | Pointer to **string** |  | [optional] 
**OwnRole** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**PlayerIdentityIds** | Pointer to **[]string** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**RevealStage** | Pointer to **int32** |  | [optional] 
**Suspects** | Pointer to [**[]VoteOption**](VoteOption.md) |  | [optional] 

## Methods

### NewFindtheminorityRoundView

`func NewFindtheminorityRoundView() *FindtheminorityRoundView`

NewFindtheminorityRoundView instantiates a new FindtheminorityRoundView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindtheminorityRoundViewWithDefaults

`func NewFindtheminorityRoundViewWithDefaults() *FindtheminorityRoundView`

NewFindtheminorityRoundViewWithDefaults instantiates a new FindtheminorityRoundView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccusationSubmittedCount

`func (o *FindtheminorityRoundView) GetAccusationSubmittedCount() int32`

GetAccusationSubmittedCount returns the AccusationSubmittedCount field if non-nil, zero value otherwise.

### GetAccusationSubmittedCountOk

`func (o *FindtheminorityRoundView) GetAccusationSubmittedCountOk() (*int32, bool)`

GetAccusationSubmittedCountOk returns a tuple with the AccusationSubmittedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccusationSubmittedCount

`func (o *FindtheminorityRoundView) SetAccusationSubmittedCount(v int32)`

SetAccusationSubmittedCount sets AccusationSubmittedCount field to given value.

### HasAccusationSubmittedCount

`func (o *FindtheminorityRoundView) HasAccusationSubmittedCount() bool`

HasAccusationSubmittedCount returns a boolean if a field has been set.

### GetAccusedAnswer

`func (o *FindtheminorityRoundView) GetAccusedAnswer() string`

GetAccusedAnswer returns the AccusedAnswer field if non-nil, zero value otherwise.

### GetAccusedAnswerOk

`func (o *FindtheminorityRoundView) GetAccusedAnswerOk() (*string, bool)`

GetAccusedAnswerOk returns a tuple with the AccusedAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccusedAnswer

`func (o *FindtheminorityRoundView) SetAccusedAnswer(v string)`

SetAccusedAnswer sets AccusedAnswer field to given value.

### HasAccusedAnswer

`func (o *FindtheminorityRoundView) HasAccusedAnswer() bool`

HasAccusedAnswer returns a boolean if a field has been set.

### GetAccusedIdentityId

`func (o *FindtheminorityRoundView) GetAccusedIdentityId() string`

GetAccusedIdentityId returns the AccusedIdentityId field if non-nil, zero value otherwise.

### GetAccusedIdentityIdOk

`func (o *FindtheminorityRoundView) GetAccusedIdentityIdOk() (*string, bool)`

GetAccusedIdentityIdOk returns a tuple with the AccusedIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccusedIdentityId

`func (o *FindtheminorityRoundView) SetAccusedIdentityId(v string)`

SetAccusedIdentityId sets AccusedIdentityId field to given value.

### HasAccusedIdentityId

`func (o *FindtheminorityRoundView) HasAccusedIdentityId() bool`

HasAccusedIdentityId returns a boolean if a field has been set.

### GetAnswerCounts

`func (o *FindtheminorityRoundView) GetAnswerCounts() map[string]int32`

GetAnswerCounts returns the AnswerCounts field if non-nil, zero value otherwise.

### GetAnswerCountsOk

`func (o *FindtheminorityRoundView) GetAnswerCountsOk() (*map[string]int32, bool)`

GetAnswerCountsOk returns a tuple with the AnswerCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerCounts

`func (o *FindtheminorityRoundView) SetAnswerCounts(v map[string]int32)`

SetAnswerCounts sets AnswerCounts field to given value.

### HasAnswerCounts

`func (o *FindtheminorityRoundView) HasAnswerCounts() bool`

HasAnswerCounts returns a boolean if a field has been set.

### GetAnswerSubmittedCount

`func (o *FindtheminorityRoundView) GetAnswerSubmittedCount() int32`

GetAnswerSubmittedCount returns the AnswerSubmittedCount field if non-nil, zero value otherwise.

### GetAnswerSubmittedCountOk

`func (o *FindtheminorityRoundView) GetAnswerSubmittedCountOk() (*int32, bool)`

GetAnswerSubmittedCountOk returns a tuple with the AnswerSubmittedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerSubmittedCount

`func (o *FindtheminorityRoundView) SetAnswerSubmittedCount(v int32)`

SetAnswerSubmittedCount sets AnswerSubmittedCount field to given value.

### HasAnswerSubmittedCount

`func (o *FindtheminorityRoundView) HasAnswerSubmittedCount() bool`

HasAnswerSubmittedCount returns a boolean if a field has been set.

### GetAnswers

`func (o *FindtheminorityRoundView) GetAnswers() map[string]string`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *FindtheminorityRoundView) GetAnswersOk() (*map[string]string, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *FindtheminorityRoundView) SetAnswers(v map[string]string)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *FindtheminorityRoundView) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetAudienceBoard

`func (o *FindtheminorityRoundView) GetAudienceBoard() VoteTallyView`

GetAudienceBoard returns the AudienceBoard field if non-nil, zero value otherwise.

### GetAudienceBoardOk

`func (o *FindtheminorityRoundView) GetAudienceBoardOk() (*VoteTallyView, bool)`

GetAudienceBoardOk returns a tuple with the AudienceBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceBoard

`func (o *FindtheminorityRoundView) SetAudienceBoard(v VoteTallyView)`

SetAudienceBoard sets AudienceBoard field to given value.

### HasAudienceBoard

`func (o *FindtheminorityRoundView) HasAudienceBoard() bool`

HasAudienceBoard returns a boolean if a field has been set.

### GetAudiencePollState

`func (o *FindtheminorityRoundView) GetAudiencePollState() string`

GetAudiencePollState returns the AudiencePollState field if non-nil, zero value otherwise.

### GetAudiencePollStateOk

`func (o *FindtheminorityRoundView) GetAudiencePollStateOk() (*string, bool)`

GetAudiencePollStateOk returns a tuple with the AudiencePollState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiencePollState

`func (o *FindtheminorityRoundView) SetAudiencePollState(v string)`

SetAudiencePollState sets AudiencePollState field to given value.

### HasAudiencePollState

`func (o *FindtheminorityRoundView) HasAudiencePollState() bool`

HasAudiencePollState returns a boolean if a field has been set.

### GetAudienceVerdictRevealed

`func (o *FindtheminorityRoundView) GetAudienceVerdictRevealed() bool`

GetAudienceVerdictRevealed returns the AudienceVerdictRevealed field if non-nil, zero value otherwise.

### GetAudienceVerdictRevealedOk

`func (o *FindtheminorityRoundView) GetAudienceVerdictRevealedOk() (*bool, bool)`

GetAudienceVerdictRevealedOk returns a tuple with the AudienceVerdictRevealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceVerdictRevealed

`func (o *FindtheminorityRoundView) SetAudienceVerdictRevealed(v bool)`

SetAudienceVerdictRevealed sets AudienceVerdictRevealed field to given value.

### HasAudienceVerdictRevealed

`func (o *FindtheminorityRoundView) HasAudienceVerdictRevealed() bool`

HasAudienceVerdictRevealed returns a boolean if a field has been set.

### GetBoardToAudience

`func (o *FindtheminorityRoundView) GetBoardToAudience() bool`

GetBoardToAudience returns the BoardToAudience field if non-nil, zero value otherwise.

### GetBoardToAudienceOk

`func (o *FindtheminorityRoundView) GetBoardToAudienceOk() (*bool, bool)`

GetBoardToAudienceOk returns a tuple with the BoardToAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardToAudience

`func (o *FindtheminorityRoundView) SetBoardToAudience(v bool)`

SetBoardToAudience sets BoardToAudience field to given value.

### HasBoardToAudience

`func (o *FindtheminorityRoundView) HasBoardToAudience() bool`

HasBoardToAudience returns a boolean if a field has been set.

### GetBoardToPlayers

`func (o *FindtheminorityRoundView) GetBoardToPlayers() bool`

GetBoardToPlayers returns the BoardToPlayers field if non-nil, zero value otherwise.

### GetBoardToPlayersOk

`func (o *FindtheminorityRoundView) GetBoardToPlayersOk() (*bool, bool)`

GetBoardToPlayersOk returns a tuple with the BoardToPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardToPlayers

`func (o *FindtheminorityRoundView) SetBoardToPlayers(v bool)`

SetBoardToPlayers sets BoardToPlayers field to given value.

### HasBoardToPlayers

`func (o *FindtheminorityRoundView) HasBoardToPlayers() bool`

HasBoardToPlayers returns a boolean if a field has been set.

### GetId

`func (o *FindtheminorityRoundView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindtheminorityRoundView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindtheminorityRoundView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindtheminorityRoundView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMinorityOption

`func (o *FindtheminorityRoundView) GetMinorityOption() string`

GetMinorityOption returns the MinorityOption field if non-nil, zero value otherwise.

### GetMinorityOptionOk

`func (o *FindtheminorityRoundView) GetMinorityOptionOk() (*string, bool)`

GetMinorityOptionOk returns a tuple with the MinorityOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorityOption

`func (o *FindtheminorityRoundView) SetMinorityOption(v string)`

SetMinorityOption sets MinorityOption field to given value.

### HasMinorityOption

`func (o *FindtheminorityRoundView) HasMinorityOption() bool`

HasMinorityOption returns a boolean if a field has been set.

### GetNumber

`func (o *FindtheminorityRoundView) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *FindtheminorityRoundView) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *FindtheminorityRoundView) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *FindtheminorityRoundView) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetOptions

`func (o *FindtheminorityRoundView) GetOptions() []VoteOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FindtheminorityRoundView) GetOptionsOk() (*[]VoteOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FindtheminorityRoundView) SetOptions(v []VoteOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FindtheminorityRoundView) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOutcomeKind

`func (o *FindtheminorityRoundView) GetOutcomeKind() string`

GetOutcomeKind returns the OutcomeKind field if non-nil, zero value otherwise.

### GetOutcomeKindOk

`func (o *FindtheminorityRoundView) GetOutcomeKindOk() (*string, bool)`

GetOutcomeKindOk returns a tuple with the OutcomeKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeKind

`func (o *FindtheminorityRoundView) SetOutcomeKind(v string)`

SetOutcomeKind sets OutcomeKind field to given value.

### HasOutcomeKind

`func (o *FindtheminorityRoundView) HasOutcomeKind() bool`

HasOutcomeKind returns a boolean if a field has been set.

### GetOwnAccusation

`func (o *FindtheminorityRoundView) GetOwnAccusation() string`

GetOwnAccusation returns the OwnAccusation field if non-nil, zero value otherwise.

### GetOwnAccusationOk

`func (o *FindtheminorityRoundView) GetOwnAccusationOk() (*string, bool)`

GetOwnAccusationOk returns a tuple with the OwnAccusation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnAccusation

`func (o *FindtheminorityRoundView) SetOwnAccusation(v string)`

SetOwnAccusation sets OwnAccusation field to given value.

### HasOwnAccusation

`func (o *FindtheminorityRoundView) HasOwnAccusation() bool`

HasOwnAccusation returns a boolean if a field has been set.

### GetOwnAnswer

`func (o *FindtheminorityRoundView) GetOwnAnswer() string`

GetOwnAnswer returns the OwnAnswer field if non-nil, zero value otherwise.

### GetOwnAnswerOk

`func (o *FindtheminorityRoundView) GetOwnAnswerOk() (*string, bool)`

GetOwnAnswerOk returns a tuple with the OwnAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnAnswer

`func (o *FindtheminorityRoundView) SetOwnAnswer(v string)`

SetOwnAnswer sets OwnAnswer field to given value.

### HasOwnAnswer

`func (o *FindtheminorityRoundView) HasOwnAnswer() bool`

HasOwnAnswer returns a boolean if a field has been set.

### GetOwnAudienceVote

`func (o *FindtheminorityRoundView) GetOwnAudienceVote() string`

GetOwnAudienceVote returns the OwnAudienceVote field if non-nil, zero value otherwise.

### GetOwnAudienceVoteOk

`func (o *FindtheminorityRoundView) GetOwnAudienceVoteOk() (*string, bool)`

GetOwnAudienceVoteOk returns a tuple with the OwnAudienceVote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnAudienceVote

`func (o *FindtheminorityRoundView) SetOwnAudienceVote(v string)`

SetOwnAudienceVote sets OwnAudienceVote field to given value.

### HasOwnAudienceVote

`func (o *FindtheminorityRoundView) HasOwnAudienceVote() bool`

HasOwnAudienceVote returns a boolean if a field has been set.

### GetOwnRole

`func (o *FindtheminorityRoundView) GetOwnRole() string`

GetOwnRole returns the OwnRole field if non-nil, zero value otherwise.

### GetOwnRoleOk

`func (o *FindtheminorityRoundView) GetOwnRoleOk() (*string, bool)`

GetOwnRoleOk returns a tuple with the OwnRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnRole

`func (o *FindtheminorityRoundView) SetOwnRole(v string)`

SetOwnRole sets OwnRole field to given value.

### HasOwnRole

`func (o *FindtheminorityRoundView) HasOwnRole() bool`

HasOwnRole returns a boolean if a field has been set.

### GetPhase

`func (o *FindtheminorityRoundView) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *FindtheminorityRoundView) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *FindtheminorityRoundView) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *FindtheminorityRoundView) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPlayerIdentityIds

`func (o *FindtheminorityRoundView) GetPlayerIdentityIds() []string`

GetPlayerIdentityIds returns the PlayerIdentityIds field if non-nil, zero value otherwise.

### GetPlayerIdentityIdsOk

`func (o *FindtheminorityRoundView) GetPlayerIdentityIdsOk() (*[]string, bool)`

GetPlayerIdentityIdsOk returns a tuple with the PlayerIdentityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerIdentityIds

`func (o *FindtheminorityRoundView) SetPlayerIdentityIds(v []string)`

SetPlayerIdentityIds sets PlayerIdentityIds field to given value.

### HasPlayerIdentityIds

`func (o *FindtheminorityRoundView) HasPlayerIdentityIds() bool`

HasPlayerIdentityIds returns a boolean if a field has been set.

### GetQuestion

`func (o *FindtheminorityRoundView) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *FindtheminorityRoundView) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *FindtheminorityRoundView) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *FindtheminorityRoundView) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetRevealStage

`func (o *FindtheminorityRoundView) GetRevealStage() int32`

GetRevealStage returns the RevealStage field if non-nil, zero value otherwise.

### GetRevealStageOk

`func (o *FindtheminorityRoundView) GetRevealStageOk() (*int32, bool)`

GetRevealStageOk returns a tuple with the RevealStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevealStage

`func (o *FindtheminorityRoundView) SetRevealStage(v int32)`

SetRevealStage sets RevealStage field to given value.

### HasRevealStage

`func (o *FindtheminorityRoundView) HasRevealStage() bool`

HasRevealStage returns a boolean if a field has been set.

### GetSuspects

`func (o *FindtheminorityRoundView) GetSuspects() []VoteOption`

GetSuspects returns the Suspects field if non-nil, zero value otherwise.

### GetSuspectsOk

`func (o *FindtheminorityRoundView) GetSuspectsOk() (*[]VoteOption, bool)`

GetSuspectsOk returns a tuple with the Suspects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspects

`func (o *FindtheminorityRoundView) SetSuspects(v []VoteOption)`

SetSuspects sets Suspects field to given value.

### HasSuspects

`func (o *FindtheminorityRoundView) HasSuspects() bool`

HasSuspects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


