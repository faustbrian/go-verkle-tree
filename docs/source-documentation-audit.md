# Source documentation audit

## Ownership and acceptance

Documentation-only Tier A pass at remote-main base `b9df78f61c6ab7cbd1c04342804ac57abe4183f0`. The root library owns a package-defined experimental research profile, not production cryptographic assurance or Ethereum compatibility. Direct consumers use only the root package; the seven internal packages own its state, layout, vector, encoding and backend implementation. No owned reverse dependency is declared in `modules.json`. Caller-owned stores and trusted roots remain application boundaries.

Before edits: retain all executable tokens, AST, literals, directives, public API baseline, dependencies, profile/source pins, official fixtures, corpus and decision history. Improve package/API/field contracts and non-obvious internal rationale only. Inspect all tracked Go, including private helpers and test-only race variants. Smallest checks: pinned formatting and all-source identity, eight native package renders, the existing output-bearing `ExampleSnapshot`, native documentation/spelling/local links and offline specification checks, current structural source checker, changed-file/added-content/staged secrets and one complete ordinary final diff. No cryptographic, conformance, race, fuzz, performance, 386 or release campaign.

## Initial gaps and questions

The initial structural inventory reports 198 errors: 197 undocumented internal exported fields and one missing backend package header. These are documentation gaps, not a runtime or cryptographic verdict. Existing root documentation already describes most public budgets and immutable values. Investigate admission versus authentication, point/scalar canonicality, per-engine queue serialization, dependency-owned mutable global authority, cancellation checkpoints, conservative rather than process-wide memory accounting, storage capability declarations versus adapter enforcement, atomic publication/crash windows, and ownership of returned buffers.

Unresolved security authority is preserved: pinned IPA backend production audit and complete soundness are not established; exceptional transcript challenge handling remains unresolved; admitted dependency calls cannot be interrupted; dependency globals remain authoritative; complete side-channel/native-platform qualification is absent. Existing pinned differential vectors establish only their selected corpora. Do not infer production readiness from stable API metadata. No source patch can close these questions in this pass.

## Native producer and immutable boundaries

`.golib.yaml` uses `apidiff` with `api/baseline.txt`, not a full-prose API snapshot; preserve that baseline. CI consumes Tools v2.3.0 at the exact configured source pin. Root Make has no documentation Example target; the existing single Example is selected directly. Generated Go source or Go generation directives are absent in the tracked inventory. Pinned Go/Rust fixture producers under `interoperability/` are invoked by `scripts/check-interoperability.sh` through `verification/package.mk`; those research authorities are not executed here. Test-only race flags exclude the uninstrumented allocation assertion under race or coverage instrumentation. Security decisions, source revisions and dated backend evidence remain unchanged.

## Reproducible complete inventory

Reproduce with `git ls-files -- "*.go"`; module/package boundaries come from `modules.json` and `packages.json`. Each row records source roles, native entry-point names and initial structural gaps. Test files are read for contract partitions, oracles, helpers and limits; no unchanged campaign is implied.

| Tracked file | Role / entry points | Initial gaps |
| --- | --- | --- |
| `aaa_00_storage_termination_internal_test.go` | test; TestStorageAccountingAndUpdateConversionTerminate | 0 |
| `architecture_test.go` | test; TestProductionSourceHasNoPackageInitOrOwnedGoroutines, TestProductionSourceArchitectureRespectsModuleOwnership, TestSourceArchitectureViolationDetection | 0 |
| `doc.go` | source | 0 |
| `errors.go` | source | 0 |
| `example_test.go` | test; ExampleSnapshot | 0 |
| `facade_errors_test.go` | test; TestFacadeResourceErrorMappings, TestFacadeErrorTranslationFallbacks, TestFacadeWitnessErrorTranslations | 0 |
| `facade_methods_test.go` | test; TestFacadeRootRejectsHostileAndZeroValues, TestFacadeSnapshotRejectsEveryInvalidOwnershipState, TestFacadeSnapshotPreflightsResourcesBeforeCopyingCallerSlices, TestFacadeProofRejectsEveryInvalidOwnershipState, TestFacadeUpdateProofRejectsInvalidAndBoundedInputs, TestFacadeWitnessRejectsEveryInvalidOwnershipState, TestFacadeTrustedRootMismatchPrecedesBackendVerification, TestFacadeLimitValidationChecksEveryFieldAndBoundary | 0 |
| `internal/authstate/aaa_failfast_boundaries_test.go` | test; TestFailFastTreeProofSuffixHalfGroupingProgresses, TestFailFastTreeProofMergeCopiesSortedValues, TestFailFastAggregateVerifierQueryCapacityBoundsTopology, TestFailFastStatelessWitnessStemPathLookup, TestFailFastCanonicalEmptyRootStemPath, TestFailFastEmptyRootTreeProofShape, TestFailFastClaimAccessorsPreservePresenceAndAbsence, TestFailFastClaimLimitsRejectEachInvalidField, TestFailFastAggregateVerifierQueriesEmptyRoot, TestFailFastAggregateVerifierCollectorStemQueries, TestFailFastProofMaterialRemainingBoundary, TestFailFastSnapshotLimitDisjunction, TestFailFastCopyEntriesAcceptsExactMaximumCount, TestFailFastClaimMergeObservesCancellation, TestFailFastClaimMergeSortsTwoValues, TestFailFastUpdateProofAccountsForEveryWorkingKey, TestFailFastStatelessProofKeyAccountsForNextSlot, TestFailFastStatelessDeleteScratchAccountsForWholeVector, TestFailFastStatelessScratchAccountsForEveryComponent, TestFailFastStatelessAbsentDeleteDoesNotStopLaterStemUpdate, TestFailFastStatelessAbsentDeleteDoesNotStopSameStemInsertion, TestFailFastStatelessNoOpDoesNotStopLaterStemUpdate, TestFailFastStatelessMissingInsertionClassifiesChildShape, TestFailFastUpdateProofClassifiesWholeStemTransitions, TestFailFastUpdateProofContinuesAcrossStemGroups, TestFailFastStemDepthUsesNeighborAfterMultiSuffixStem, TestFailFastMergeExistingStemIntoEmptyInsertion, TestFailFastParentChangedChildReturnsExactIndex, TestFailFastDisclosedChildClassifiesEveryPathKind | 0 |
| `internal/authstate/aggregate_queries.go` | source | 8 |
| `internal/authstate/aggregate_queries_test.go` | test; TestProofMaterialReconstructsVerifiableAggregateQueries, TestEmptyRootVerifierQueriesConsolidateSharedRootIndex, TestAggregateVerifierQueriesBoundInitialDeduplicationCapacity, TestGrowAggregateVerifierQueriesCapsAllocatedCapacity, TestAggregateVerifierQueriesTraverseEveryInternalDepth, TestAggregateQueriesConsolidateSharedOpenings, TestAggregateVerifierQueriesRejectSurplusCommitment, TestAggregateVerifierQueriesOpenDifferentStem, TestAggregateVerifierQueriesPreserveCancellation, TestAggregateVerifierQueriesRejectInvalidInputsAndResources, TestAggregateVerifierQueryCapacityBoundsTopology, TestAggregateVerifierQueryHelpersRejectInvalidState | 0 |
| `internal/authstate/claims.go` | source | 5 |
| `internal/authstate/claims_benchmark_test.go` | test; BenchmarkNewClaimSetSixteen, BenchmarkClaimSetCopySixteen | 0 |
| `internal/authstate/claims_fuzz_test.go` | test; FuzzClaimSetCanonicalization | 0 |
| `internal/authstate/claims_test.go` | test; TestClaimSetCanonicalizesAndOwnsClaims, TestClaimSetRejectsDuplicatesAndInvalidClaims, TestClaimSetRejectsProfileBeforeResources, TestClaimSetEnforcesResourcesBeforeAllocation, TestClaimSetRejectsInvalidStateContextAndLimits, TestClaimLimitsAndSetValidationHonorImplementationBoundaries, TestClaimSetCanonicalizesMergeBoundaries, TestClaimSetSupportsConcurrentImmutableReads | 0 |
| `internal/authstate/error_classification_test.go` | test; TestFacadeErrorClassifiers | 0 |
| `internal/authstate/proof_engine.go` | source | 5 |
| `internal/authstate/proof_engine_benchmark_test.go` | test; BenchmarkProofEngine | 0 |
| `internal/authstate/proof_engine_test.go` | test; TestProofEngineGeneratesDeterministicVerifiableTreeProof, TestProofEngineGeneratesVerifiableEmptyRootNonMembership, TestProofEngineProvesCompleteTopologyDeletionTransition, TestProofEngineDerivesCanonicalUpdateProofKeys, TestProofEngineUpdateProofRejectsInvalidAndBoundedInputs, TestStatelessProofKeyHelpersEnforceBoundsAndCancellation, TestStatelessSnapshotStemDepthUsesCanonicalNeighborCollisions, TestProofEngineUpdateProofCancellationBoundaries, TestProofEngineRejectsTamperedProofs, TestProofEngineRejectsInvalidInputsBeforeOpeningWork, TestProofEnginePropagatesBoundedStageFailures, TestProofEngineStatementBindingFailures, TestProofEngineVerificationPreservesMaterialAndCancellationFailures, TestProofEngineSupportsConcurrentImmutableUse, TestMatchAggregateQueriesRejectsDivergence | 0 |
| `internal/authstate/proof_material.go` | source | 9 |
| `internal/authstate/proof_material_benchmark_test.go` | test; BenchmarkSnapshotProofMaterialSixteen | 0 |
| `internal/authstate/proof_material_test.go` | test; TestSnapshotProofMaterialDerivesCanonicalClaimsAndPaths, TestSnapshotProofMaterialDerivesEmptyRootNonMembership, TestSnapshotProofMaterialRejectsInvalidInputs, TestProofMaterialLimitsRejectZeroAndExcessiveValues, TestSnapshotProofMaterialEnforcesAggregateResources, TestProofMaterialExtractionRejectsAlreadyExhaustedAccounting, TestProofMaterialExtractionAccountsSuccessfulPaths, TestSnapshotProofMaterialAccountsSecondSuffixHalf, TestSnapshotProofMaterialSupportsMaximumStemDepth, TestSnapshotProofMaterialAvoidsRedundantStemExtractions, TestSnapshotProofMaterialExtractsSuffixBoundaryHalf, TestSnapshotProofMaterialCancellationCheckpoints, TestProofMaterialIsDeterministicOwnedAndConcurrent, TestProofMaterialRejectsInvalidReceiversAndCopyContexts, TestProofMaterialHelpers | 0 |
| `internal/authstate/root_container_test.go` | test; TestSnapshotRootContainerBindsProfileAndCommitment, TestSnapshotRootContainerRepresentsEmptyRoot, TestTransitionRootContainersBindExactSnapshots, TestRootContainersRejectInvalidStateAndCancellation | 0 |
| `internal/authstate/snapshot.go` | source | 6 |
| `internal/authstate/snapshot_benchmark_test.go` | test; BenchmarkSnapshotGet, BenchmarkSnapshotApply | 0 |
| `internal/authstate/snapshot_fuzz_test.go` | test; FuzzSnapshotTransitions | 0 |
| `internal/authstate/snapshot_test.go` | test; TestApplyMatchesPinnedIndependentPreAndPostRoots, TestSnapshotAppliesCanonicalImmutableTransitions, TestSnapshotCanonicalOrderDoesNotDependOnInputOrder, TestSnapshotMergeCoversOrderedBoundaryCases, TestSnapshotTransitionsMatchIndependentStateModel, TestSnapshotExhaustiveReducedDomainMatchesIndependentModel, TestSnapshotRejectsInvalidTransitionsAtomically, TestSnapshotEnforcesLimitsBeforeAmplifiedWork, TestSnapshotRejectsInvalidStateInputsAndContexts, TestSnapshotPropagatesCommittedTreeFailures, TestSnapshotStopsAtCancellationBoundaries, TestCanonicalSortHelpersPreserveExactAndStableOrder, TestSnapshotSupportsConcurrentImmutableUse, TestSnapshotCopiesCanonicalEntriesWithinBounds | 0 |
| `internal/authstate/stateless_update.go` | source | 8 |
| `internal/authstate/stateless_update_benchmark_test.go` | test; BenchmarkStatelessUpdaterTransitions | 0 |
| `internal/authstate/stateless_update_fuzz_test.go` | test; FuzzStatelessUpdaterMatchesStatefulTransition | 0 |
| `internal/authstate/stateless_update_interop_test.go` | test; TestStatelessUpdaterMatchesPinnedRustRebuiltTransitions | 0 |
| `internal/authstate/stateless_update_test.go` | test; TestStatelessUpdaterDerivesPinnedPostStateRoot, TestStatelessUpdaterDeletesWithoutChangingTopology, TestStatelessUpdaterTreatsAuthenticatedAbsentDeletesAsNoOps, TestStatelessUpdaterAppliesMixedSetsAndDeletes, TestStatelessUpdaterCombinesExistingStemMutationWithInsertion, TestStatelessUpdaterDeletesLastStemWithCompleteTopologyProof, TestStatelessUpdaterCollapsesUnaryInternalPathAfterStemDeletion, TestStatelessUpdaterCollapsesEveryUnaryAncestorAfterStemDeletion, TestStatelessUpdaterCollapsesBatchDeletedSiblingStems, TestStatelessStemRetentionHonorsCancellationAndOrdering, TestStatelessStemRetentionStopsAtNextStem, TestStatelessUpdaterInsertsOneAuthenticatedMissingStem, TestStatelessUpdaterBuildsCanonicalSubtreeForMissingSiblingStems, TestStatelessUpdaterSplitsAuthenticatedDifferentStem, TestStatelessUpdaterBuildsDeepestDifferentStemChain, TestStatelessUpdaterHandlesDeepSharedPathsAndBothSuffixHalves, TestStatelessUpdaterSupportsPresentZeroAndNoOpValues, TestStatelessUpdaterRejectsInvalidAndIncompleteRequests, TestNewStatelessUpdaterFromProofEngineReusesOpeningSettings, TestStatelessUpdaterEnforcesEveryResource, TestStatelessUpdaterAccountsForProofAndPropagationScratch, TestStatelessUpdateProofCountAndScratchBoundaries, TestStatelessProofIndexesHonorCancellation, TestStatelessProofIndexesIncludeRootWithoutPathCommitments, TestStatelessUpdaterHonorsCancellationAndConcurrentUse, TestStatelessUpdaterHonorsTopologyInsertionCancellation, TestStatelessUpdateInternalFailureBoundaries, TestStatelessTopologyDisclosureFailureBoundaries, TestStatelessDisclosedParentFailureBoundaries, TestStatelessInsertedTopologyInternalFailureBoundaries | 0 |
| `internal/authstate/stateless_witness.go` | source | 13 |
| `internal/authstate/stateless_witness_benchmark_test.go` | test; BenchmarkEncodeStatelessWitness, BenchmarkDecodeStatelessWitness, BenchmarkApplyStatelessWitness | 0 |
| `internal/authstate/stateless_witness_fuzz_test.go` | test; FuzzDecodeStatelessWitness, FuzzApplyStatelessWitness | 0 |
| `internal/authstate/stateless_witness_test.go` | test; TestStatelessWitnessCanonicalRoundTrip, TestStatelessWitnessCanonicalDeleteRoundTrip, TestStatelessWitnessAcceptsCompleteTopologyDeletionProof, TestStatelessWitnessTopologyClaimFailureBoundaries, TestStatelessWitnessDecoderRejectsShortEmbeddedProof, TestStatelessWitnessUpdateKeysAreStrictlyIncreasing, TestStatelessWitnessRelationBytes, TestStatelessUpdaterVerifiesWitnessPostStateRoot, TestStatelessWitnessRejectsInvalidConstructionAndAccess, TestStatelessWitnessAcceptsExactMaximumLimits, TestStatelessWitnessCancellationBoundaries, TestStatelessWitnessEncodingRejectsMalformedAndExhaustiveLimits, TestStatelessWitnessErrorsAndUpdateAccessors | 0 |
| `internal/authstate/tree_proof.go` | source | 9 |
| `internal/authstate/tree_proof_benchmark_test.go` | test; BenchmarkNewTreeProofSixteen, BenchmarkTreeProofCopySixteen | 0 |
| `internal/authstate/tree_proof_encoding.go` | source | 17 |
| `internal/authstate/tree_proof_encoding_benchmark_test.go` | test; BenchmarkEncodeTreeProof, BenchmarkDecodeTreeProof, BenchmarkDecodeTreeProofWrongLength | 0 |
| `internal/authstate/tree_proof_encoding_fuzz_test.go` | test; FuzzDecodeTreeProof, FuzzVerifyTreeProof | 0 |
| `internal/authstate/tree_proof_encoding_test.go` | test; TestTreeProofBytesUsesCanonicalProfileBoundEncoding, TestTreeProofEncodingUsesExactPreV1Layout, TestDecodeTreeProofRoundTripsCanonicalEncoding, TestTreeProofEncodingRoundTripsExplicitEmptyPathCommitment, TestTreeProofEmptySuffixRequiresOnlyAbsenceClaims, TestTreeProofEncodingLimitBoundaries, TestTreeProofEncodingRejectsInvalidStateContextAndResources, TestDecodeTreeProofAcceptsEveryCanonicalClaimAndStemPathKind, TestDecodeTreeProofDistinguishesTruncatedHeaderFromDeclaredLength, TestDecodeTreeProofPreservesMultipleRecordStrides, TestDecodeTreeProofRejectsNonCanonicalRecordOrdering, TestDecodeTreeProofAcceptsMaximumPathLength, TestDecodeTreeProofRejectsMalformedCanonicalEncoding, TestTreeProofIdentityOpeningRequiresCryptographicVerification, TestDecodeTreeProofEnforcesResourcesBeforeCryptographicDecoding, TestDecodeTreeProofRejectsInvalidLimitsAndHonorsCancellation | 0 |
| `internal/authstate/tree_proof_fuzz_test.go` | test; FuzzTreeProofCanonicalization | 0 |
| `internal/authstate/tree_proof_markers_test.go` | test; TestDerivePathMarkersCoalescesSuffixHalves, TestMatchExpectedCommitmentsRequiresEveryIdentityPermission | 0 |
| `internal/authstate/tree_proof_sort_test.go` | test; TestTreeProofIndexSortMatchesStableReference | 0 |
| `internal/authstate/tree_proof_test.go` | test; TestTreeProofCanonicalizesAndBindsComponents, TestTreeProofRequiresExactCanonicalPathSet, TestTreeProofAcceptsCanonicalEmptyRootNonMembership, TestTreeProofRejectsNonCanonicalEmptyRootShapes, TestTreeProofPathDerivationStopsDuringClaimGrouping, TestStemPathRejectsInvalidMetadata, TestPathCommitmentValidatesAndOwnsPath, TestTreeProofRejectsInvalidComponentsAndTopology, TestTreeProofRejectsConflictingSharedTopology, TestTreeProofDeduplicatesSharedSuffixCommitmentPaths, TestTreeProofInternalBoundariesFailClosed, TestTreeProofEnforcesEveryResourceLimit, TestTreeProofTemporaryAccountingIsExact, TestTreeProofCancellationAndInvalidReceiverBehavior, TestTreeProofHelpersPreserveBoundariesAndStableOrder, TestTreeProofSupportsConcurrentImmutableReads | 0 |
| `internal/backend/aggregate_opening.go` | source | 18 |
| `internal/backend/aggregate_opening_benchmark_test.go` | test; BenchmarkAggregateOpeningEngineOpenSharedVectors | 0 |
| `internal/backend/aggregate_opening_test.go` | test; TestAggregateOpeningEngineMatchesPinnedRustProof, TestBoundAggregateOpeningRejectsStatementReplayForZeroVector, TestAggregateOpeningEngineBoundsConcurrentDependencyWorkers, TestAggregateOpeningGateBoundsAndCancelsQueuedWork, TestAggregateOpeningEngineRejectsInvalidSetup, TestAggregateOpeningEngineRejectsMutatedAuxiliaryGenerator, TestAggregateOpeningEnginePreflightsSetupResources, TestAggregateOpeningOperationsRejectInvalidInputsAndResources, TestAggregateOpeningSupportsOneAuthenticatedZeroOpening, TestAggregateOpeningOperationsRejectMalformedQueries, TestAggregateOpeningReusesCommittedVectorPreparation, TestAggregateOpeningPreparationCapacityBoundsOneVector, TestAggregateOpeningBindingLookupSkipsInvalidQueriesAndRequiresExactAnchor, TestAggregateOpeningOperationsPropagateBackendAndCancellationFailures, TestAggregateOpeningProofNativeConversionRejectsCorruption | 0 |
| `internal/backend/commitment_decode.go` | source | 5 |
| `internal/backend/commitment_decode_test.go` | test; TestDecodeVectorCommitmentReturnsStrictCanonicalValue | 0 |
| `internal/backend/commitment_engine.go` | source | 10 |
| `internal/backend/commitment_engine_fuzz_test.go` | test; FuzzCommitmentEngine | 0 |
| `internal/backend/commitment_engine_interop_test.go` | test; TestCommitmentEngineMatchesPinnedRustVectors, TestCommitmentEngineSparseUpdateMatchesPinnedRustVector | 0 |
| `internal/backend/commitment_engine_test.go` | test; TestCommitmentEngineCommitsZeroVectorAsInternalIdentity, TestCommitmentEngineConcurrentCommitsAreDeterministic, TestCommitmentEngineChecksCancellationDuringCommitWork, TestCommitmentEngineUpdatesValidatedVectorPositions, TestCommitmentEngineOwnsUpdatesBeforeArithmetic, TestCommitmentEngineRejectsGeneratorSubstitution, TestCommitmentEngineRejectsInvalidAndExhaustedLimits, TestCommitmentEngineRejectsInvalidScalarsAndCommitBudgets, TestCommitmentEngineRejectsNilAndCancelledContexts, TestVectorCommitmentZeroValueRejectsUse, TestVectorCommitmentDeduplicationKey | 0 |
| `internal/backend/commitment_to_scalar_interop_test.go` | test; TestRustVerkleCommitmentToScalarVectors | 0 |
| `internal/backend/commitment_to_scalar_test.go` | test; TestCommitmentToScalarMatchesPinnedGeneratorVector | 0 |
| `internal/backend/commitment_update_test.go` | test; TestCommitmentEngineUpdateCommitmentHandlesIdentityAndNoOp, TestCommitmentEngineEqualScalarUpdateDoesNotAllocate, TestCommitmentEngineReportsSparseUpdateCapacity, TestCommitmentEngineUpdateCommitmentAcceptsEveryVectorPosition, TestCommitmentEngineUpdateCommitmentRejectsInvalidInputs, TestCommitmentEngineUpdateCommitmentEnforcesResourcesAndCancellation | 0 |
| `internal/backend/encoding.go` | source | 0 |
| `internal/backend/encoding_benchmark_test.go` | test; BenchmarkDecodeCommitmentCanonical, BenchmarkCommitVectorSparse, BenchmarkCommitVectorDense, BenchmarkUpdateCommitmentSparse, BenchmarkDecodeCommitmentIdentity, BenchmarkEncodeCommitment, BenchmarkCommitmentToScalar, BenchmarkDecodeScalarCanonical, BenchmarkDecodeScalarNonCanonical, BenchmarkEncodeScalar | 0 |
| `internal/backend/encoding_fuzz_test.go` | test; FuzzDecodeCommitment, FuzzDecodeScalar | 0 |
| `internal/backend/encoding_interop_test.go` | test; TestRustVerkleEncodingVectors, TestRustVerkleGeneratorSet, TestRustVerkleMultiProof, TestRustVerkleMultiProofRejectsMutations | 0 |
| `internal/backend/encoding_test.go` | test; TestCommitmentEncodingRoundTrip, TestDecodeCommitmentRejectsIdentity, TestDecodeOpeningProofPointRejectsWrongLength, TestDecodeCommitmentRejectsMalformedEncodings, TestScalarEncodingRoundTrip, TestDecodeScalarRejectsMalformedEncodings | 0 |
| `internal/backend/opening_proof.go` | source | 6 |
| `internal/backend/opening_proof_benchmark_test.go` | test; BenchmarkDecodeOpeningProofCanonical, BenchmarkDecodeOpeningProofWrongLength | 0 |
| `internal/backend/opening_proof_fuzz_test.go` | test; FuzzDecodeOpeningProof | 0 |
| `internal/backend/opening_proof_test.go` | test; TestOpeningProofMatchesPinnedRustFixture, TestOpeningProofRejectsMalformedEncodings, TestOpeningProofAcceptsCanonicalIdentityElements, TestOpeningProofEnforcesResourcesBeforeDecoding, TestOpeningProofRejectsInvalidLimitsStateAndContext, TestOpeningProofValidatesEveryEncodedPoint | 0 |
| `internal/backend/race_disabled_test.go` | test; build !race | 0 |
| `internal/backend/race_enabled_test.go` | test; build race | 0 |
| `internal/backend/root_container.go` | source | 5 |
| `internal/backend/root_container_benchmark_test.go` | test; BenchmarkDecodeRootCommitment, BenchmarkDecodeRootEmpty, BenchmarkDecodeRootWrongLength | 0 |
| `internal/backend/root_container_fuzz_test.go` | test; FuzzDecodeRoot | 0 |
| `internal/backend/root_container_test.go` | test; TestRootContainerCanonicalRoundTrip, TestRootContainerEncodesEmptyRootExplicitly, TestRootContainerRejectsMalformedEncodings, TestRootContainerRejectsProfileBeforePointBudget, TestRootContainerEnforcesResources, TestRootContainerRejectsInvalidStateContextAndLimits, TestInvalidRootRejectsCommitmentAccess | 0 |
| `internal/committedtree/aggregate_queries.go` | source | 10 |
| `internal/committedtree/aggregate_queries_test.go` | test; TestAggregateProverQueriesAreCanonicalAndComplete, TestStemVectorsUseZeroVectorForAbsentC1Half, TestAggregateProverQueriesRejectInvalidInputsAndResources, TestAggregateProverQueriesBoundScratchByDistinctStems, TestAggregateProverQueriesAccountExactDistinctStemPrefixCapacity, TestAggregateProverQueriesStopCapacityAtFirstDistinctStemByte, TestGrowAggregateProverQueriesCapsAllocatedCapacity, TestAggregateProverQueriesPreserveCancellation, TestAggregateProverQueriesTraverseInternalAndDifferentStems, TestAggregateQueryCollectorRejectsCorruptTrees, TestAggregateQueryHelpersRejectInvalidState, TestAggregateQuerySortHelpers, TestBuildPreparedRejectsFinalizedCountMismatch | 0 |
| `internal/committedtree/bounded_context_test.go` | test | 0 |
| `internal/committedtree/proof_path.go` | source | 14 |
| `internal/committedtree/proof_path_test.go` | test; TestProofPathExtractsCanonicalMembershipAndAbsenceMaterial, TestProofPathIntoReusesCallerScratch, TestProofPathEnforcesLimitsAndCallerOwnership, TestProofPathRejectsInvalidUseAndHonorsCancellation, TestProofPathLimitsAndCorruptTreesFailClosed, TestProofPathAcceptsMaximumDepth | 0 |
| `internal/committedtree/storage_decoding.go` | source | 8 |
| `internal/committedtree/storage_decoding_fuzz_test.go` | test; FuzzDecodeStorageNode | 0 |
| `internal/committedtree/storage_decoding_test.go` | test; TestDecodeStorageNodeRoundTripsCanonicalImage, TestDecodeStorageNodeRejectsHostileEnvelopeBeforePointDecoding, TestDecodeStorageNodeRejectsInvalidInputsAndBudgets, TestDecodeStorageNodeRejectsNonCanonicalNodeBodies, TestDecodedStorageNodeAccessorsFailClosedAndRemainCancellable, TestStorageDecodingArithmeticAndBoundariesAreExact, TestDecodeStorageNodeReportsEveryDecodedPoint | 0 |
| `internal/committedtree/storage_encoding.go` | source | 8 |
| `internal/committedtree/storage_encoding_test.go` | test; TestStorageImageOwnsCanonicalContentAddressedNodes, TestStorageImageIsDeterministicAcrossInputOrder, TestStorageImageEncodesEmptyTreeAsOneAddressableNode, TestStorageImageAcceptsExactEmptyTreeResourceBounds, TestStorageImageRejectsInvalidStateLimitsAndCancellation, TestStorageEncodingErrorsAndZeroValuesFailClosed, TestStorageEncodingLimitsRejectCountsBeyondSupportedBoundary, TestStorageImageAccountsExactMultiNodeBytes, TestStorageNodeSizeRejectsEveryCorruptArenaInvariant, TestStorageImageRejectsEveryCorruptTopologyInvariant, TestStorageTreeValidationRejectsEachInvalidStateIndependently, TestStorageTopologyRejectsEachRootInvariantIndependently, TestStorageImageAcceptsNestedCanonicalTopology, TestStorageImageAcceptsMaximumCanonicalDepth, TestValidateStorageSubtreeRejectsDefensiveInvalidState, TestStorageEncodingHelpersFailClosedAndRemainCancellable, TestSortStorageNodesOrdersBoundariesAndPreservesEqualIDs, TestStorageImageRemainsCancellableAcrossEveryEncodingPhase | 0 |
| `internal/committedtree/tree.go` | source | 13 |
| `internal/committedtree/tree_benchmark_test.go` | test; BenchmarkBuildFourEntries, BenchmarkBuilderTopologyPreservingUpdate, BenchmarkProofPath, BenchmarkStorageImage | 0 |
| `internal/committedtree/tree_fuzz_test.go` | test; FuzzBuildDeterministic, FuzzBuilderUpdateMatchesFullRebuild | 0 |
| `internal/committedtree/tree_internal_test.go` | test; TestNewBuilderRejectsInvalidInputs, TestBuilderZeroAndCorruptValuesRejectUse, TestBuilderUpdateRejectsEachInvalidBuilderInvariant, TestBuilderUpdateRejectsInvalidTreeAndPreparation, TestBuilderUpdatePropagatesIncrementalCommitmentFailure, TestBuilderUpdateCancellationBoundaries, TestSparseStemUpdateRejectsEveryInvalidCommitmentBoundary, TestBuilderUpdateRejectsInvalidSparseStemResult, TestTopologyRebuilderPropagatesCommitmentAndTopologyFailures, TestTopologyRebuilderRejectsCorruptRemovedCommitment, TestTopologyLookupAndRebuildCancellationBoundaries, TestTopologyLookupsEnforceCanonicalDepthAndEdges, TestBuilderUpdateReusesIdenticalTree, TestBuilderUpdateSkipsUnchangedInternalBranch, TestBuilderUpdateRejectsInvalidUnchangedChildCommitment, TestSparseVectorUpdateChunksAndFallsBackDeterministically, TestBuilderPropagatesPreparationFailure, TestPrepareBuildPropagatesEveryCancellationBoundary, TestPrepareBuildRejectsIntermediateScratchBudget, TestExactConfiguredBoundariesAreAccepted, TestPrepareBuildBoundsAggregateCryptographicWork, TestContextAwareConstructionHelpersCancel, TestBuildPreparedPropagatesCommitmentFailure, TestTreeBuilderPropagatesCommitAndMappingFailures, TestTreeBuilderChecksContext | 0 |
| `internal/committedtree/tree_interop_test.go` | test; TestBuildMatchesPinnedRustTreeRoots | 0 |
| `internal/committedtree/tree_test.go` | test; TestBuildMatchesPinnedIndependentTreeRoot, TestBuilderReusesImmutableCommitmentEngineConcurrently, TestBuilderUpdateRecommitsOnlyChangedExistingStemPath, TestBuilderUpdateMatchesRebuildAcrossValueAndTopologyChanges, TestBuilderUpdateReusesCommitmentsAcrossTopologyChanges, TestBuildEmptyTreeRetainsIdentityRoot, TestBuildCanonicalizesAndOwnsEntries, TestBuildDistinguishesPresentZeroFromAbsence, TestBuildRetainsCanonicalCollisionTopology, TestBuildRejectsDuplicateKeys, TestBuildRejectsInvalidLimits, TestBuildRejectsExhaustedResourcesBeforeCommitmentWork, TestBuildRejectsInvalidCommitmentLimits, TestBuildRejectsNilAndCancelledContexts, TestTreeZeroAndCorruptValuesRejectUse, TestConstructionInvariantHelpersFailClosed | 0 |
| `internal/leafvector/vector.go` | source | 5 |
| `internal/leafvector/vector_interop_test.go` | test; TestRustVerkleLeafVectors | 0 |
| `internal/leafvector/vector_test.go` | test; TestEncodePresentValue, TestPresentZeroDiffersFromAbsent, TestSuffixPlacement, TestEverySuffixHasContiguousInRangeIndices, TestEncodeStem, TestStemCommitmentInputPositions | 0 |
| `internal/profile/profile.go` | source | 0 |
| `internal/profile/profile_test.go` | test; TestProfileIsExactAndRejectsOtherValues | 0 |
| `internal/statemodel/doc.go` | source | 0 |
| `internal/statemodel/errors.go` | source | 3 |
| `internal/statemodel/state.go` | source | 3 |
| `internal/statemodel/state_fuzz_test.go` | test; FuzzSnapshotTransitions | 0 |
| `internal/statemodel/state_test.go` | test; TestSnapshotAppliesCanonicalImmutableTransitions, TestSnapshotCanonicalOrderDoesNotDependOnBatchOrder, TestSnapshotMergeCoversOrderedBoundaryCases, TestSnapshotRejectsInvalidTransitionsAtomically, TestSnapshotEnforcesOperationBudgetsBeforeAllocation, TestSnapshotRejectsInvalidStateAndContext, TestSnapshotAcceptsExactTemporaryBudget, TestSnapshotAccountsForOwnedUpdateAndSortScratch, TestSnapshotCountsReplacementAndDeletionIndependently, TestSnapshotStopsDuringBoundedLoops, TestEmptyApplyReturnsEquivalentSnapshot, TestSortUpdatesPreservesEqualKeyOrder, TestSortUpdatesHonorsEveryCancellationBoundary | 0 |
| `internal/treelayout/doc.go` | source | 0 |
| `internal/treelayout/errors.go` | source | 3 |
| `internal/treelayout/layout.go` | source | 7 |
| `internal/treelayout/layout_fuzz_test.go` | test; FuzzLayoutTransitions | 0 |
| `internal/treelayout/layout_interop_test.go` | test; TestLayoutMatchesPinnedRustPathHints | 0 |
| `internal/treelayout/layout_test.go` | test; TestBuildCreatesCanonicalTopologyAndClassifiesPaths, TestBuildIsDeterministicAndDoesNotMutateCallerInput, TestLayoutInsertAndDeleteReturnCanonicalImmutableValues, TestBuildRejectsDuplicatesAndInvalidInputs, TestLayoutRejectsCancelledOperations, TestBuildEnforcesEveryResourceBudget, TestLayoutReportsDeterministicConstructionBytes, TestLayoutFailsClosedForCorruptOwnedTopology, TestLayoutChecksCancellationDuringBoundedWork, TestLayoutMutationOperationsPropagateResourceAndCancellationErrors, TestSortStemsMatchesCanonicalOrderForSmallPermutations, TestSortStemsHonorsEveryCancellationBoundary | 0 |
| `main_test.go` | test; TestMain | 0 |
| `profile.go` | source | 0 |
| `profile_test.go` | test; TestBandersnatchIPA256V0, TestZeroProfileIsUnsupported | 0 |
| `proof.go` | source | 0 |
| `proof_test.go` | test; TestPublicProofEngineGeneratesCanonicalVerifiableProofs, TestPublicProofEngineVerifiesTrustedRootAndKeySet, TestPublicProofEngineProvesCanonicalEmptyRootNonMembership, TestPublicProofEngineRejectsTamperingAndInvalidUse | 0 |
| `public_benchmark_test.go` | test; BenchmarkPublicSnapshotOperations, BenchmarkPublicProofOperations, BenchmarkPublicStatelessWitnessOperations, BenchmarkPublicStatelessEngineInitialization, BenchmarkPublicStatelessTopologyTransitions | 0 |
| `root.go` | source | 0 |
| `snapshot.go` | source | 0 |
| `snapshot_encoding.go` | source | 0 |
| `snapshot_encoding_fuzz_test.go` | test; FuzzDecodeSnapshot | 0 |
| `snapshot_encoding_internal_test.go` | test; TestSnapshotEncodingCancellationAndErrorClassification, TestSnapshotEncodingLimitCeilings | 0 |
| `snapshot_encoding_test.go` | test; TestPublicSnapshotEncodingIsCanonicalAndSelfAuthenticating, TestDecodeSnapshotRejectsMalformedAndExhaustedInput, TestSnapshotEncodingRejectsInvalidUseAndOwnsEmptyState, TestSnapshotEncodingPreflightsEveryAllocation, BenchmarkEncodeSnapshotTwoEntries, BenchmarkDecodeSnapshotTwoEntries | 0 |
| `snapshot_test.go` | test; TestPublicSnapshotSupportsImmutableAuthenticatedTransitions, TestPublicSnapshotIsDeterministicAndFailsClosed | 0 |
| `stateless_transition_fuzz_test.go` | test; FuzzPublicStatelessTransitionMatchesStatefulSnapshot | 0 |
| `storage.go` | source | 0 |
| `storage_audit.go` | source | 0 |
| `storage_audit_benchmark_test.go` | test; BenchmarkAuditStorageCurrentAndRetainedSnapshots | 0 |
| `storage_audit_fuzz_test.go` | test; FuzzStorageAuditInventorySequence | 0 |
| `storage_audit_internal_test.go` | test; TestAuditStorageFindsOnlyNodesOutsideCurrentAndRetainedSnapshots, TestAuditStorageRejectsInventoryThatOmitsReachableNode, TestAuditStorageRejectsInvalidInputsBeforeOpeningStore, TestStorageAuditLimitsRejectEveryInvalidField, TestAuditStoragePreservesLifecycleAndReachableReadFailures, TestAuditStorageRejectsMalformedPublicationSets, TestAuditStorageRejectsMalformedAndExhaustedInventoryPages, TestAuditStorageSupportsEmptyPublicationSetAndBoundsUnreachableNodes, TestStorageAuditZeroValueAndCancellationFailClosed, TestAuditStorageRejectsCancellationAfterEmptyInventoryRead, TestStorageAuditInternalBoundariesFailClosed, TestStorageAuditRejectsHiddenPageCapacityAndPreflightsResultGrowth, TestStorageAuditAccountingFunctionsAreExact | 0 |
| `storage_audit_width_internal_test.go` | test; TestStorageAuditResultCapacityPreserves32BitBounds | 0 |
| `storage_crash_test.go` | test; TestStorageCommitCrashMatrixAndRetry, TestStorageMaintenanceCrashRetryAndPinnedAuditView, TestStorageRecoveryCrashIsAtomicAndRetryable | 0 |
| `storage_internal_test.go` | test; TestStorageFacadeValuesAndErrorsFailClosed, TestStorageLimitsRejectEachInvalidFieldAndAcceptBoundaries, TestSnapshotCommitTranslatesEveryStorageResource, TestSnapshotCommitRejectsInvalidPreviousAndAcceptsValueStore, TestSnapshotCommitRemainsCancellableAcrossEncodingAndCopying, TestTranslateStorageEncodingErrorDefaultResource | 0 |
| `storage_maintenance.go` | source | 0 |
| `storage_maintenance_benchmark_test.go` | test; BenchmarkMaintainStorageDropRetainedAndPrune, BenchmarkRecoverStoragePreserveRetainedAndDeleteOrphan | 0 |
| `storage_maintenance_failures_internal_test.go` | test; TestMaintainStoragePreservesPlanningAndCancellationFailures, TestMaintainStorageEnforcesPlanningResourceBudgets, TestCopyRequestedRetainedChecksBeforeAllocatingAndEachElement, TestMaintainStorageChargesEveryLivePublicationAndReachabilitySlot, TestMaintainStorageChargesDroppedReachabilityUntilFinalRequest, TestMaintainStorageChecksFinalRequestAllocationBudget, TestNormalizeRetainedPublicationsChecksBeforeCopying, TestNormalizeRetainedPublicationsAcceptsExactCountAndCapacityLimit, TestNormalizeRetainedPublicationsChargesCurrentPublication, TestStorePublicationMergeSortCanonicalizesMultipleValuesAndRejectsDuplicates, TestStorePublicationMergeSortObservesEveryCancellationPhase, TestMaintenanceInventoryRejectsEveryMalformedBoundary, TestMaintenanceInventoryEnforcesExactPageMemoryAndDeletionBoundaries, TestMaintenanceInventoryRejectsDuplicateIDsWithinAndAcrossPages, TestStoreMaintenanceAccessorsPreserveValidationCancellationAndOwnership, TestMaintenanceOrderingHelpersCoverEqualAndGreaterNodeIDs | 0 |
| `storage_maintenance_fuzz_test.go` | test; FuzzStorageMaintenanceInventorySequence | 0 |
| `storage_maintenance_internal_test.go` | test; TestRecoverStoragePreservesPublicationsAndRemovesInterruptedWrites, TestRecoverStorageSupportsAnUnpublishedNodeOnlyStore, TestRecoverStorageFailsAtomicallyForInvalidOrChangingState, TestRecoverStorageRejectsInvalidInputsBeforeOpening, TestMaintainStorageAtomicallyAppliesRetentionAndPruning, TestMaintainStorageLeavesACompleteAuditablePostMaintenanceStore, TestMaintainStorageReadsOneIsolatedPublicationSet, TestMaintainStorageRejectsInvalidInputsAndCapabilitiesBeforeOpening, TestMaintainStorageRejectsUnboundProfileNamespaceBeforeOpening, TestMaintainStorageRejectsInvalidRetentionSetsWithoutApplying, TestMaintainStorageValidatesAndCopiesRequestedRetentionBeforeOpening, TestMaintainStoragePreservesLifecycleAndAtomicFailureSemantics, TestMaintainStorageRejectsIncompleteInventoryAndDeletionBudget, TestMaintainStorageCanonicalizesRetentionAndVerifiesDroppedSnapshots, TestMaintainStorageSupportsNoCurrentPublicationAndRejectsMalformedPages, TestMaintainStorageStopsAfterCloseCancellation, TestStoreMaintenanceAndResultZeroValuesRejectUse | 0 |
| `storage_read.go` | source | 0 |
| `storage_read_benchmark_test.go` | test; BenchmarkLoadSnapshotFourEntries | 0 |
| `storage_read_fuzz_test.go` | test; FuzzLoadSnapshotRootNode | 0 |
| `storage_read_internal_test.go` | test; TestStoragePublicationAndReadLimitsFailClosed, TestLoadSnapshotRejectsFacadeAndReaderLifecycleFailures, TestLoadSnapshotEnforcesEveryReadResource, TestLoadSnapshotPassesExactRemainingReadBounds, TestLoadSnapshotAccountsPointDecodesAcrossNodes, TestLoadSnapshotRemainsCancellableAcrossReadsAndReconstruction, TestLoadSnapshotRejectsAdapterThatExceedsReadBound, TestLoadSnapshotTemporaryBudgetStopsEachAllocationPhase, TestLoadSnapshotIsDeterministicAcrossConcurrentReadViews, TestLoadSnapshotRejectsCanonicalTopologyAndCommitmentSubstitution, TestStorageReadHelpersCoverOverflowCancellationAndTranslations | 0 |
| `storage_test.go` | test; TestSnapshotCommitPublishesCompleteOwnedCanonicalBatch, TestSnapshotCommitWithoutPreviousRootPreservesExpectation, TestSnapshotCommitRejectsInvalidInputsAndCapabilitiesBeforeStoreCall, TestSnapshotCommitWrapsStoreFailureAndPreservesSnapshot, TestLoadSnapshotReconstructsPublishedCanonicalState, TestLoadSnapshotReconstructsEmptyRootWithoutPointDecoding, TestLoadSnapshotDistinguishesMissingCorruptAndStoreFailures | 0 |
| `witness.go` | source | 0 |
| `witness_test.go` | test; TestPublicStatelessEngineInitializesStandalone, TestPublicStatelessWitnessRoundTripAndApplication, TestPublicStatelessWitnessAppliesTopologyChangingSets, TestPublicStatelessWitnessAppliesDeletionWithoutTopologyChange, TestPublicStatelessWitnessAppliesInsertionToEmptyRoot, TestPublicProofEngineBuildsAndAppliesTopologyDeletionWitness, TestPublicStatelessWitnessRejectsInvalidUseAndTampering | 0 |

Initial total: 138 tracked Go files; one module; eight source packages.
Five new comment-only canonical headers complete the final 143-file inventory:

| Added tracked file | Role | Initial gaps |
| --- | --- | --- |
| `internal/authstate/doc.go` | canonical package contract | new |
| `internal/backend/doc.go` | canonical package contract | new |
| `internal/committedtree/doc.go` | canonical package contract | new |
| `internal/leafvector/doc.go` | canonical package contract | new |
| `internal/profile/doc.go` | canonical package contract | new |

The sole module is `github.com/faustbrian/go-verkle-tree` at `.`. Package
directories are `.`, `internal/authstate`, `internal/backend`,
`internal/committedtree`, `internal/leafvector`, `internal/profile`,
`internal/statemodel`, and `internal/treelayout`; their catalog import paths are
the module path plus the directory. Only the root is public.

## Source-traced contracts and disposition

| Owner / direct consumer | Audited contract and adjacent evidence |
| --- | --- |
| Root facade / application callers | Fixed arrays copy keys and values; constructors copy slices while reading them synchronously. Duplicate initial/batch keys fail; empty snapshot batches preserve state; present zero differs from Delete. Zero opaque values reject use. Root and proof decoding establish representation, not root authority. `VerifyForKeys` requires the exact root and distinct key set; `ApplyForRoot` compares trusted pre-root before proof arithmetic. Facade ownership, error and expectation tests cover these distinctions. |
| `internal/profile`, `internal/leafvector` / committed tree | Descriptor validation compares every fixed field. Values map into two half vectors with a low-half presence marker; stem is 31-byte little-endian, not a key hash. Scalar inputs are bounded below the field modulus. Frozen profile and pinned leaf-vector differential tests provide the rationale, not production assurance. |
| `internal/statemodel`, `internal/treelayout` / tests and topology oracle | Immutable sorted arrays, duplicate rejection, minimal radix paths, absent/different/present stem outcomes; budgeted copying/merge sorting and checked counts precede arena construction. Oracle and fuzz inputs remain test-only, including separate same-named transition fuzz targets. |
| `internal/backend` / tree and proof engines | Exact width-256 generators are checked by ordered canonical-byte digest and auxiliary generator identity. Commitment updates decode old/new canonical scalars and reject duplicate positions; callers must authenticate Old first. Standalone commitments reject identity, root containers encode an explicit empty kind, and raw IPA proof elements permit canonical identity. Scalars use strict little-endian decoding; representation decoding does not verify openings. Pinned point/scalar/Rust corpora, zero-vector and sparse-update tests substantiate only their recorded partitions. |
| `internal/backend` / concurrent proof calls | CPU-derived dependency workers are admitted by budget, not configured by MaxWorkers. Setup and admitted backend work cannot be interrupted. One active slot and bounded cancellable waiters are per engine; copies/reused engines share the gate, separately constructed engines do not. Distinct commitment copies protect against dependency in-place batch normalization. Public dependency globals remain authoritative. Gate, cancellation, setup-substitution and statement-replay tests explain wrapper invariants without closing backend blockers. |
| `internal/committedtree` / authenticated state and loader | Tree arenas own canonical entries/nodes/edges. Update reuses unchanged commitments and authenticated scalar deltas; capacity zero falls back to full commit. ProofPathInto may alias caller scratch. Prover queries own vectors, subsequently borrowed by the backend. Canonical storage SHA-256 addresses bind full profile envelopes; local decoding is not complete tree authentication. Existing topology, ownership, pinned interop and bounded-context tests cover these partitions. |
| `internal/authstate` / public proofs and witnesses | Exact sorted claims, stem outcomes and path commitments reject omitted/surplus topology. Public verifier queries are reconstructed independently of private vectors. Statement digest binds root, claims, paths, commitments, indices and evaluations; a nonzero anchor prevents all-zero statement replay. Stateless updates verify old claims before arithmetic. Emptying a stem requires suffix and ancestor completeness probes before canonical collapse; absence is not inferred from omitted data. Existing proof/witness/tamper/completeness tests and stateful oracle remain unchanged. |
| Storage facade / caller adapters | Capability bits are adapter assertions, not a verified persistence implementation. Commit performs one callback, no retry; an error cannot resolve an ambiguous adapter publication outcome. Read/audit views are closed once with the operation context, including cancellation; a Close error invalidates results. ReadNode transfers stable owned bytes and must bound allocation/I/O before return. Loader checks address, ordering, path/topology, canonical rebuilt root and root-node address. |
| Audit/maintenance / caller adapters | Inventory is strictly ordered, bounded in length/capacity and terminal-page work; unreachable bytes are reported without authentication. Per-publication read budgets multiply by MaxPublications. Audit is not deletion authority. Maintenance closes its isolated view before callback and requires atomic comparison of the full observed publication set to prevent a stale-plan deletion window. Adapter must preserve earlier views. Recovery handles node-only unpublished debris, not missing/corrupt published nodes. Crash, lifecycle, changing-view, inventory, capacity and 32-bit termination sources supply the rationale. |

### Variants, examples, markers and provenance

All 98 original test sources were inspected for their contract partitions, helper
ownership, saved hostile inputs and pinned oracles. Test-only `race` / `!race`
backend flags are intentional variants; no production build-tag variant exists.
The one `ExampleSnapshot` has an Output assertion. Fuzz and benchmark entry names
are listed above; no performance claims or new campaigns are made. Root TestMain
uses goleak as part of the existing Example execution. No generated Go or producer
directive is present. Reference corpus production remains in pinned external
Go/Rust scripts and manifests; those scripts are not run or rewritten. Security,
profile, source catalog, frozen fixture bytes and two typed decision records and
history are unchanged.

### Questions retained, not repaired

Production backend/soundness, transcript exceptional challenges, mutable globals,
non-interruptible dependency work, independent-engine aggregate worker demand,
side channels and native-platform qualification remain limited exactly as the
existing security authority states. The ordered-generator digest is an integrity
comparison, not independent setup provenance. No new practical transcript
preimage or security exploit is claimed. Store capability honesty, durable
publication and recovery from ambiguous callback errors cannot be proven without
an actual caller adapter; none is implemented here. Source provides no stronger
process-memory or hard cancellation guarantee, so comments do not invent one.

Current maintained README/reference/boundary/API-audit prose incorrectly denied
commitment reuse and conflated storage capability declarations with guarantees.
Only those present-contract statements are corrected; dated benchmark results,
backend evidence, profile policy, fixture provenance and API checksum stay intact.

### Selected verification

Passed with Go 1.27.2 on darwin/arm64: all 143 files formatted; all original
138 files preserve executable tokens/literals/AST and all 81 annotations exactly,
with no optional-semicolon normalization needed. Five new files contain only
package declarations and comments. All 1,139 comment groups rendered as HTML;
eight native `go doc -all` package renders passed. The single actual
`ExampleSnapshot` passed under its existing TestMain. Tools v2.3.0 native docs,
spelling and local links passed (107 links, zero errors, two existing exclusions),
as did offline specification validation (one module, two decisions). The first
docs attempt rejected an inherited bare executable name; formatting `rustc`
and the adjacent executable names as inline code preserves the prerequisite
and existing spelling rules. No dictionary or policy was changed.

Configured apidiff baseline is preserved, not regenerated. Current structural
source check reports zero findings over all 143 staged source paths. No broad
runtime campaign, external oracle generation, online authority refresh, hosted
CI, production cryptographic qualification or release is claimed. Ordinary final review refined deletion-stem
completeness wording (no Set in the
stem) and clarified pre-deduplication path budgets. A narrow final refresh passed
exact AST/raw tokens/literals and three annotations in the two affected files,
rendered their 70 comment groups, and refreshed the owning authstate GoDoc.
Earlier unchanged renders, Example, documentation and specification results are
retained under their original inputs, not described as repeated runs. Both
sessions' exact three caches, helpers and outputs were removed with absence
verified. Ordinary complete-diff review found no remaining material
documentation findings. The changed-current full-file secret scan passed; final
added-content and staged scans are reported in the local delivery handoff.
