// Code generated by "stringer -type=APIErrorCode -trimprefix=Err api-errors.go"; DO NOT EDIT.

package cmd

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrNone-0]
	_ = x[ErrAccessDenied-1]
	_ = x[ErrBadDigest-2]
	_ = x[ErrEntityTooSmall-3]
	_ = x[ErrEntityTooLarge-4]
	_ = x[ErrPolicyTooLarge-5]
	_ = x[ErrIncompleteBody-6]
	_ = x[ErrInternalError-7]
	_ = x[ErrInvalidAccessKeyID-8]
	_ = x[ErrAccessKeyDisabled-9]
	_ = x[ErrInvalidArgument-10]
	_ = x[ErrInvalidBucketName-11]
	_ = x[ErrInvalidDigest-12]
	_ = x[ErrInvalidRange-13]
	_ = x[ErrInvalidRangePartNumber-14]
	_ = x[ErrInvalidCopyPartRange-15]
	_ = x[ErrInvalidCopyPartRangeSource-16]
	_ = x[ErrInvalidMaxKeys-17]
	_ = x[ErrInvalidEncodingMethod-18]
	_ = x[ErrInvalidMaxUploads-19]
	_ = x[ErrInvalidMaxParts-20]
	_ = x[ErrInvalidPartNumberMarker-21]
	_ = x[ErrInvalidPartNumber-22]
	_ = x[ErrInvalidRequestBody-23]
	_ = x[ErrInvalidCopySource-24]
	_ = x[ErrInvalidMetadataDirective-25]
	_ = x[ErrInvalidCopyDest-26]
	_ = x[ErrInvalidPolicyDocument-27]
	_ = x[ErrInvalidObjectState-28]
	_ = x[ErrMalformedXML-29]
	_ = x[ErrMissingContentLength-30]
	_ = x[ErrMissingContentMD5-31]
	_ = x[ErrMissingRequestBodyError-32]
	_ = x[ErrMissingSecurityHeader-33]
	_ = x[ErrNoSuchBucket-34]
	_ = x[ErrNoSuchBucketPolicy-35]
	_ = x[ErrNoSuchBucketLifecycle-36]
	_ = x[ErrNoSuchLifecycleConfiguration-37]
	_ = x[ErrInvalidLifecycleWithObjectLock-38]
	_ = x[ErrNoSuchBucketSSEConfig-39]
	_ = x[ErrNoSuchCORSConfiguration-40]
	_ = x[ErrNoSuchWebsiteConfiguration-41]
	_ = x[ErrReplicationConfigurationNotFoundError-42]
	_ = x[ErrRemoteDestinationNotFoundError-43]
	_ = x[ErrReplicationDestinationMissingLock-44]
	_ = x[ErrRemoteTargetNotFoundError-45]
	_ = x[ErrReplicationRemoteConnectionError-46]
	_ = x[ErrReplicationBandwidthLimitError-47]
	_ = x[ErrBucketRemoteIdenticalToSource-48]
	_ = x[ErrBucketRemoteAlreadyExists-49]
	_ = x[ErrBucketRemoteLabelInUse-50]
	_ = x[ErrBucketRemoteArnTypeInvalid-51]
	_ = x[ErrBucketRemoteArnInvalid-52]
	_ = x[ErrBucketRemoteRemoveDisallowed-53]
	_ = x[ErrRemoteTargetNotVersionedError-54]
	_ = x[ErrReplicationSourceNotVersionedError-55]
	_ = x[ErrReplicationNeedsVersioningError-56]
	_ = x[ErrReplicationBucketNeedsVersioningError-57]
	_ = x[ErrReplicationDenyEditError-58]
	_ = x[ErrRemoteTargetDenyAddError-59]
	_ = x[ErrReplicationNoExistingObjects-60]
	_ = x[ErrReplicationValidationError-61]
	_ = x[ErrReplicationPermissionCheckError-62]
	_ = x[ErrObjectRestoreAlreadyInProgress-63]
	_ = x[ErrNoSuchKey-64]
	_ = x[ErrNoSuchUpload-65]
	_ = x[ErrInvalidVersionID-66]
	_ = x[ErrNoSuchVersion-67]
	_ = x[ErrNotImplemented-68]
	_ = x[ErrPreconditionFailed-69]
	_ = x[ErrRequestTimeTooSkewed-70]
	_ = x[ErrSignatureDoesNotMatch-71]
	_ = x[ErrMethodNotAllowed-72]
	_ = x[ErrInvalidPart-73]
	_ = x[ErrInvalidPartOrder-74]
	_ = x[ErrMissingPart-75]
	_ = x[ErrAuthorizationHeaderMalformed-76]
	_ = x[ErrMalformedPOSTRequest-77]
	_ = x[ErrPOSTFileRequired-78]
	_ = x[ErrSignatureVersionNotSupported-79]
	_ = x[ErrBucketNotEmpty-80]
	_ = x[ErrAllAccessDisabled-81]
	_ = x[ErrPolicyInvalidVersion-82]
	_ = x[ErrMissingFields-83]
	_ = x[ErrMissingCredTag-84]
	_ = x[ErrCredMalformed-85]
	_ = x[ErrInvalidRegion-86]
	_ = x[ErrInvalidServiceS3-87]
	_ = x[ErrInvalidServiceSTS-88]
	_ = x[ErrInvalidRequestVersion-89]
	_ = x[ErrMissingSignTag-90]
	_ = x[ErrMissingSignHeadersTag-91]
	_ = x[ErrMalformedDate-92]
	_ = x[ErrMalformedPresignedDate-93]
	_ = x[ErrMalformedCredentialDate-94]
	_ = x[ErrMalformedExpires-95]
	_ = x[ErrNegativeExpires-96]
	_ = x[ErrAuthHeaderEmpty-97]
	_ = x[ErrExpiredPresignRequest-98]
	_ = x[ErrRequestNotReadyYet-99]
	_ = x[ErrUnsignedHeaders-100]
	_ = x[ErrMissingDateHeader-101]
	_ = x[ErrInvalidQuerySignatureAlgo-102]
	_ = x[ErrInvalidQueryParams-103]
	_ = x[ErrBucketAlreadyOwnedByYou-104]
	_ = x[ErrInvalidDuration-105]
	_ = x[ErrBucketAlreadyExists-106]
	_ = x[ErrMetadataTooLarge-107]
	_ = x[ErrUnsupportedMetadata-108]
	_ = x[ErrUnsupportedHostHeader-109]
	_ = x[ErrMaximumExpires-110]
	_ = x[ErrSlowDownRead-111]
	_ = x[ErrSlowDownWrite-112]
	_ = x[ErrMaxVersionsExceeded-113]
	_ = x[ErrInvalidPrefixMarker-114]
	_ = x[ErrBadRequest-115]
	_ = x[ErrKeyTooLongError-116]
	_ = x[ErrInvalidBucketObjectLockConfiguration-117]
	_ = x[ErrObjectLockConfigurationNotFound-118]
	_ = x[ErrObjectLockConfigurationNotAllowed-119]
	_ = x[ErrNoSuchObjectLockConfiguration-120]
	_ = x[ErrObjectLocked-121]
	_ = x[ErrInvalidRetentionDate-122]
	_ = x[ErrPastObjectLockRetainDate-123]
	_ = x[ErrUnknownWORMModeDirective-124]
	_ = x[ErrBucketTaggingNotFound-125]
	_ = x[ErrObjectLockInvalidHeaders-126]
	_ = x[ErrInvalidTagDirective-127]
	_ = x[ErrPolicyAlreadyAttached-128]
	_ = x[ErrPolicyNotAttached-129]
	_ = x[ErrExcessData-130]
	_ = x[ErrInvalidEncryptionMethod-131]
	_ = x[ErrInvalidEncryptionKeyID-132]
	_ = x[ErrInsecureSSECustomerRequest-133]
	_ = x[ErrSSEMultipartEncrypted-134]
	_ = x[ErrSSEEncryptedObject-135]
	_ = x[ErrInvalidEncryptionParameters-136]
	_ = x[ErrInvalidEncryptionParametersSSEC-137]
	_ = x[ErrInvalidSSECustomerAlgorithm-138]
	_ = x[ErrInvalidSSECustomerKey-139]
	_ = x[ErrMissingSSECustomerKey-140]
	_ = x[ErrMissingSSECustomerKeyMD5-141]
	_ = x[ErrSSECustomerKeyMD5Mismatch-142]
	_ = x[ErrInvalidSSECustomerParameters-143]
	_ = x[ErrIncompatibleEncryptionMethod-144]
	_ = x[ErrKMSNotConfigured-145]
	_ = x[ErrKMSKeyNotFoundException-146]
	_ = x[ErrKMSDefaultKeyAlreadyConfigured-147]
	_ = x[ErrNoAccessKey-148]
	_ = x[ErrInvalidToken-149]
	_ = x[ErrEventNotification-150]
	_ = x[ErrARNNotification-151]
	_ = x[ErrRegionNotification-152]
	_ = x[ErrOverlappingFilterNotification-153]
	_ = x[ErrFilterNameInvalid-154]
	_ = x[ErrFilterNamePrefix-155]
	_ = x[ErrFilterNameSuffix-156]
	_ = x[ErrFilterValueInvalid-157]
	_ = x[ErrOverlappingConfigs-158]
	_ = x[ErrUnsupportedNotification-159]
	_ = x[ErrContentSHA256Mismatch-160]
	_ = x[ErrContentChecksumMismatch-161]
	_ = x[ErrStorageFull-162]
	_ = x[ErrRequestBodyParse-163]
	_ = x[ErrObjectExistsAsDirectory-164]
	_ = x[ErrInvalidObjectName-165]
	_ = x[ErrInvalidObjectNamePrefixSlash-166]
	_ = x[ErrInvalidResourceName-167]
	_ = x[ErrInvalidLifecycleQueryParameter-168]
	_ = x[ErrServerNotInitialized-169]
	_ = x[ErrBucketMetadataNotInitialized-170]
	_ = x[ErrRequestTimedout-171]
	_ = x[ErrClientDisconnected-172]
	_ = x[ErrTooManyRequests-173]
	_ = x[ErrInvalidRequest-174]
	_ = x[ErrTransitionStorageClassNotFoundError-175]
	_ = x[ErrInvalidStorageClass-176]
	_ = x[ErrBackendDown-177]
	_ = x[ErrMalformedJSON-178]
	_ = x[ErrAdminNoSuchUser-179]
	_ = x[ErrAdminNoSuchUserLDAPWarn-180]
	_ = x[ErrAdminLDAPExpectedLoginName-181]
	_ = x[ErrAdminNoSuchGroup-182]
	_ = x[ErrAdminGroupNotEmpty-183]
	_ = x[ErrAdminGroupDisabled-184]
	_ = x[ErrAdminInvalidGroupName-185]
	_ = x[ErrAdminNoSuchJob-186]
	_ = x[ErrAdminNoSuchPolicy-187]
	_ = x[ErrAdminPolicyChangeAlreadyApplied-188]
	_ = x[ErrAdminInvalidArgument-189]
	_ = x[ErrAdminInvalidAccessKey-190]
	_ = x[ErrAdminInvalidSecretKey-191]
	_ = x[ErrAdminConfigNoQuorum-192]
	_ = x[ErrAdminConfigTooLarge-193]
	_ = x[ErrAdminConfigBadJSON-194]
	_ = x[ErrAdminNoSuchConfigTarget-195]
	_ = x[ErrAdminConfigEnvOverridden-196]
	_ = x[ErrAdminConfigDuplicateKeys-197]
	_ = x[ErrAdminConfigInvalidIDPType-198]
	_ = x[ErrAdminConfigLDAPNonDefaultConfigName-199]
	_ = x[ErrAdminConfigLDAPValidation-200]
	_ = x[ErrAdminConfigIDPCfgNameAlreadyExists-201]
	_ = x[ErrAdminConfigIDPCfgNameDoesNotExist-202]
	_ = x[ErrInsecureClientRequest-203]
	_ = x[ErrObjectTampered-204]
	_ = x[ErrAdminLDAPNotEnabled-205]
	_ = x[ErrSiteReplicationInvalidRequest-206]
	_ = x[ErrSiteReplicationPeerResp-207]
	_ = x[ErrSiteReplicationBackendIssue-208]
	_ = x[ErrSiteReplicationServiceAccountError-209]
	_ = x[ErrSiteReplicationBucketConfigError-210]
	_ = x[ErrSiteReplicationBucketMetaError-211]
	_ = x[ErrSiteReplicationIAMError-212]
	_ = x[ErrSiteReplicationConfigMissing-213]
	_ = x[ErrSiteReplicationIAMConfigMismatch-214]
	_ = x[ErrAdminRebalanceAlreadyStarted-215]
	_ = x[ErrAdminRebalanceNotStarted-216]
	_ = x[ErrAdminBucketQuotaExceeded-217]
	_ = x[ErrAdminNoSuchQuotaConfiguration-218]
	_ = x[ErrHealNotImplemented-219]
	_ = x[ErrHealNoSuchProcess-220]
	_ = x[ErrHealInvalidClientToken-221]
	_ = x[ErrHealMissingBucket-222]
	_ = x[ErrHealAlreadyRunning-223]
	_ = x[ErrHealOverlappingPaths-224]
	_ = x[ErrIncorrectContinuationToken-225]
	_ = x[ErrEmptyRequestBody-226]
	_ = x[ErrUnsupportedFunction-227]
	_ = x[ErrInvalidExpressionType-228]
	_ = x[ErrBusy-229]
	_ = x[ErrUnauthorizedAccess-230]
	_ = x[ErrExpressionTooLong-231]
	_ = x[ErrIllegalSQLFunctionArgument-232]
	_ = x[ErrInvalidKeyPath-233]
	_ = x[ErrInvalidCompressionFormat-234]
	_ = x[ErrInvalidFileHeaderInfo-235]
	_ = x[ErrInvalidJSONType-236]
	_ = x[ErrInvalidQuoteFields-237]
	_ = x[ErrInvalidRequestParameter-238]
	_ = x[ErrInvalidDataType-239]
	_ = x[ErrInvalidTextEncoding-240]
	_ = x[ErrInvalidDataSource-241]
	_ = x[ErrInvalidTableAlias-242]
	_ = x[ErrMissingRequiredParameter-243]
	_ = x[ErrObjectSerializationConflict-244]
	_ = x[ErrUnsupportedSQLOperation-245]
	_ = x[ErrUnsupportedSQLStructure-246]
	_ = x[ErrUnsupportedSyntax-247]
	_ = x[ErrUnsupportedRangeHeader-248]
	_ = x[ErrLexerInvalidChar-249]
	_ = x[ErrLexerInvalidOperator-250]
	_ = x[ErrLexerInvalidLiteral-251]
	_ = x[ErrLexerInvalidIONLiteral-252]
	_ = x[ErrParseExpectedDatePart-253]
	_ = x[ErrParseExpectedKeyword-254]
	_ = x[ErrParseExpectedTokenType-255]
	_ = x[ErrParseExpected2TokenTypes-256]
	_ = x[ErrParseExpectedNumber-257]
	_ = x[ErrParseExpectedRightParenBuiltinFunctionCall-258]
	_ = x[ErrParseExpectedTypeName-259]
	_ = x[ErrParseExpectedWhenClause-260]
	_ = x[ErrParseUnsupportedToken-261]
	_ = x[ErrParseUnsupportedLiteralsGroupBy-262]
	_ = x[ErrParseExpectedMember-263]
	_ = x[ErrParseUnsupportedSelect-264]
	_ = x[ErrParseUnsupportedCase-265]
	_ = x[ErrParseUnsupportedCaseClause-266]
	_ = x[ErrParseUnsupportedAlias-267]
	_ = x[ErrParseUnsupportedSyntax-268]
	_ = x[ErrParseUnknownOperator-269]
	_ = x[ErrParseMissingIdentAfterAt-270]
	_ = x[ErrParseUnexpectedOperator-271]
	_ = x[ErrParseUnexpectedTerm-272]
	_ = x[ErrParseUnexpectedToken-273]
	_ = x[ErrParseUnexpectedKeyword-274]
	_ = x[ErrParseExpectedExpression-275]
	_ = x[ErrParseExpectedLeftParenAfterCast-276]
	_ = x[ErrParseExpectedLeftParenValueConstructor-277]
	_ = x[ErrParseExpectedLeftParenBuiltinFunctionCall-278]
	_ = x[ErrParseExpectedArgumentDelimiter-279]
	_ = x[ErrParseCastArity-280]
	_ = x[ErrParseInvalidTypeParam-281]
	_ = x[ErrParseEmptySelect-282]
	_ = x[ErrParseSelectMissingFrom-283]
	_ = x[ErrParseExpectedIdentForGroupName-284]
	_ = x[ErrParseExpectedIdentForAlias-285]
	_ = x[ErrParseUnsupportedCallWithStar-286]
	_ = x[ErrParseNonUnaryAggregateFunctionCall-287]
	_ = x[ErrParseMalformedJoin-288]
	_ = x[ErrParseExpectedIdentForAt-289]
	_ = x[ErrParseAsteriskIsNotAloneInSelectList-290]
	_ = x[ErrParseCannotMixSqbAndWildcardInSelectList-291]
	_ = x[ErrParseInvalidContextForWildcardInSelectList-292]
	_ = x[ErrIncorrectSQLFunctionArgumentType-293]
	_ = x[ErrValueParseFailure-294]
	_ = x[ErrEvaluatorInvalidArguments-295]
	_ = x[ErrIntegerOverflow-296]
	_ = x[ErrLikeInvalidInputs-297]
	_ = x[ErrCastFailed-298]
	_ = x[ErrInvalidCast-299]
	_ = x[ErrEvaluatorInvalidTimestampFormatPattern-300]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternSymbolForParsing-301]
	_ = x[ErrEvaluatorTimestampFormatPatternDuplicateFields-302]
	_ = x[ErrEvaluatorTimestampFormatPatternHourClockAmPmMismatch-303]
	_ = x[ErrEvaluatorUnterminatedTimestampFormatPatternToken-304]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternToken-305]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternSymbol-306]
	_ = x[ErrEvaluatorBindingDoesNotExist-307]
	_ = x[ErrMissingHeaders-308]
	_ = x[ErrInvalidColumnIndex-309]
	_ = x[ErrAdminConfigNotificationTargetsFailed-310]
	_ = x[ErrAdminProfilerNotEnabled-311]
	_ = x[ErrInvalidDecompressedSize-312]
	_ = x[ErrAddUserInvalidArgument-313]
	_ = x[ErrAddUserValidUTF-314]
	_ = x[ErrAdminResourceInvalidArgument-315]
	_ = x[ErrAdminAccountNotEligible-316]
	_ = x[ErrAccountNotEligible-317]
	_ = x[ErrAdminServiceAccountNotFound-318]
	_ = x[ErrPostPolicyConditionInvalidFormat-319]
	_ = x[ErrInvalidChecksum-320]
	_ = x[ErrLambdaARNInvalid-321]
	_ = x[ErrLambdaARNNotFound-322]
	_ = x[ErrInvalidAttributeName-323]
	_ = x[ErrAdminNoAccessKey-324]
	_ = x[ErrAdminNoSecretKey-325]
	_ = x[ErrIAMNotInitialized-326]
	_ = x[apiErrCodeEnd-327]
}

const _APIErrorCode_name = "NoneAccessDeniedBadDigestEntityTooSmallEntityTooLargePolicyTooLargeIncompleteBodyInternalErrorInvalidAccessKeyIDAccessKeyDisabledInvalidArgumentInvalidBucketNameInvalidDigestInvalidRangeInvalidRangePartNumberInvalidCopyPartRangeInvalidCopyPartRangeSourceInvalidMaxKeysInvalidEncodingMethodInvalidMaxUploadsInvalidMaxPartsInvalidPartNumberMarkerInvalidPartNumberInvalidRequestBodyInvalidCopySourceInvalidMetadataDirectiveInvalidCopyDestInvalidPolicyDocumentInvalidObjectStateMalformedXMLMissingContentLengthMissingContentMD5MissingRequestBodyErrorMissingSecurityHeaderNoSuchBucketNoSuchBucketPolicyNoSuchBucketLifecycleNoSuchLifecycleConfigurationInvalidLifecycleWithObjectLockNoSuchBucketSSEConfigNoSuchCORSConfigurationNoSuchWebsiteConfigurationReplicationConfigurationNotFoundErrorRemoteDestinationNotFoundErrorReplicationDestinationMissingLockRemoteTargetNotFoundErrorReplicationRemoteConnectionErrorReplicationBandwidthLimitErrorBucketRemoteIdenticalToSourceBucketRemoteAlreadyExistsBucketRemoteLabelInUseBucketRemoteArnTypeInvalidBucketRemoteArnInvalidBucketRemoteRemoveDisallowedRemoteTargetNotVersionedErrorReplicationSourceNotVersionedErrorReplicationNeedsVersioningErrorReplicationBucketNeedsVersioningErrorReplicationDenyEditErrorRemoteTargetDenyAddErrorReplicationNoExistingObjectsReplicationValidationErrorReplicationPermissionCheckErrorObjectRestoreAlreadyInProgressNoSuchKeyNoSuchUploadInvalidVersionIDNoSuchVersionNotImplementedPreconditionFailedRequestTimeTooSkewedSignatureDoesNotMatchMethodNotAllowedInvalidPartInvalidPartOrderMissingPartAuthorizationHeaderMalformedMalformedPOSTRequestPOSTFileRequiredSignatureVersionNotSupportedBucketNotEmptyAllAccessDisabledPolicyInvalidVersionMissingFieldsMissingCredTagCredMalformedInvalidRegionInvalidServiceS3InvalidServiceSTSInvalidRequestVersionMissingSignTagMissingSignHeadersTagMalformedDateMalformedPresignedDateMalformedCredentialDateMalformedExpiresNegativeExpiresAuthHeaderEmptyExpiredPresignRequestRequestNotReadyYetUnsignedHeadersMissingDateHeaderInvalidQuerySignatureAlgoInvalidQueryParamsBucketAlreadyOwnedByYouInvalidDurationBucketAlreadyExistsMetadataTooLargeUnsupportedMetadataUnsupportedHostHeaderMaximumExpiresSlowDownReadSlowDownWriteMaxVersionsExceededInvalidPrefixMarkerBadRequestKeyTooLongErrorInvalidBucketObjectLockConfigurationObjectLockConfigurationNotFoundObjectLockConfigurationNotAllowedNoSuchObjectLockConfigurationObjectLockedInvalidRetentionDatePastObjectLockRetainDateUnknownWORMModeDirectiveBucketTaggingNotFoundObjectLockInvalidHeadersInvalidTagDirectivePolicyAlreadyAttachedPolicyNotAttachedExcessDataInvalidEncryptionMethodInvalidEncryptionKeyIDInsecureSSECustomerRequestSSEMultipartEncryptedSSEEncryptedObjectInvalidEncryptionParametersInvalidEncryptionParametersSSECInvalidSSECustomerAlgorithmInvalidSSECustomerKeyMissingSSECustomerKeyMissingSSECustomerKeyMD5SSECustomerKeyMD5MismatchInvalidSSECustomerParametersIncompatibleEncryptionMethodKMSNotConfiguredKMSKeyNotFoundExceptionKMSDefaultKeyAlreadyConfiguredNoAccessKeyInvalidTokenEventNotificationARNNotificationRegionNotificationOverlappingFilterNotificationFilterNameInvalidFilterNamePrefixFilterNameSuffixFilterValueInvalidOverlappingConfigsUnsupportedNotificationContentSHA256MismatchContentChecksumMismatchStorageFullRequestBodyParseObjectExistsAsDirectoryInvalidObjectNameInvalidObjectNamePrefixSlashInvalidResourceNameInvalidLifecycleQueryParameterServerNotInitializedBucketMetadataNotInitializedRequestTimedoutClientDisconnectedTooManyRequestsInvalidRequestTransitionStorageClassNotFoundErrorInvalidStorageClassBackendDownMalformedJSONAdminNoSuchUserAdminNoSuchUserLDAPWarnAdminLDAPExpectedLoginNameAdminNoSuchGroupAdminGroupNotEmptyAdminGroupDisabledAdminInvalidGroupNameAdminNoSuchJobAdminNoSuchPolicyAdminPolicyChangeAlreadyAppliedAdminInvalidArgumentAdminInvalidAccessKeyAdminInvalidSecretKeyAdminConfigNoQuorumAdminConfigTooLargeAdminConfigBadJSONAdminNoSuchConfigTargetAdminConfigEnvOverriddenAdminConfigDuplicateKeysAdminConfigInvalidIDPTypeAdminConfigLDAPNonDefaultConfigNameAdminConfigLDAPValidationAdminConfigIDPCfgNameAlreadyExistsAdminConfigIDPCfgNameDoesNotExistInsecureClientRequestObjectTamperedAdminLDAPNotEnabledSiteReplicationInvalidRequestSiteReplicationPeerRespSiteReplicationBackendIssueSiteReplicationServiceAccountErrorSiteReplicationBucketConfigErrorSiteReplicationBucketMetaErrorSiteReplicationIAMErrorSiteReplicationConfigMissingSiteReplicationIAMConfigMismatchAdminRebalanceAlreadyStartedAdminRebalanceNotStartedAdminBucketQuotaExceededAdminNoSuchQuotaConfigurationHealNotImplementedHealNoSuchProcessHealInvalidClientTokenHealMissingBucketHealAlreadyRunningHealOverlappingPathsIncorrectContinuationTokenEmptyRequestBodyUnsupportedFunctionInvalidExpressionTypeBusyUnauthorizedAccessExpressionTooLongIllegalSQLFunctionArgumentInvalidKeyPathInvalidCompressionFormatInvalidFileHeaderInfoInvalidJSONTypeInvalidQuoteFieldsInvalidRequestParameterInvalidDataTypeInvalidTextEncodingInvalidDataSourceInvalidTableAliasMissingRequiredParameterObjectSerializationConflictUnsupportedSQLOperationUnsupportedSQLStructureUnsupportedSyntaxUnsupportedRangeHeaderLexerInvalidCharLexerInvalidOperatorLexerInvalidLiteralLexerInvalidIONLiteralParseExpectedDatePartParseExpectedKeywordParseExpectedTokenTypeParseExpected2TokenTypesParseExpectedNumberParseExpectedRightParenBuiltinFunctionCallParseExpectedTypeNameParseExpectedWhenClauseParseUnsupportedTokenParseUnsupportedLiteralsGroupByParseExpectedMemberParseUnsupportedSelectParseUnsupportedCaseParseUnsupportedCaseClauseParseUnsupportedAliasParseUnsupportedSyntaxParseUnknownOperatorParseMissingIdentAfterAtParseUnexpectedOperatorParseUnexpectedTermParseUnexpectedTokenParseUnexpectedKeywordParseExpectedExpressionParseExpectedLeftParenAfterCastParseExpectedLeftParenValueConstructorParseExpectedLeftParenBuiltinFunctionCallParseExpectedArgumentDelimiterParseCastArityParseInvalidTypeParamParseEmptySelectParseSelectMissingFromParseExpectedIdentForGroupNameParseExpectedIdentForAliasParseUnsupportedCallWithStarParseNonUnaryAggregateFunctionCallParseMalformedJoinParseExpectedIdentForAtParseAsteriskIsNotAloneInSelectListParseCannotMixSqbAndWildcardInSelectListParseInvalidContextForWildcardInSelectListIncorrectSQLFunctionArgumentTypeValueParseFailureEvaluatorInvalidArgumentsIntegerOverflowLikeInvalidInputsCastFailedInvalidCastEvaluatorInvalidTimestampFormatPatternEvaluatorInvalidTimestampFormatPatternSymbolForParsingEvaluatorTimestampFormatPatternDuplicateFieldsEvaluatorTimestampFormatPatternHourClockAmPmMismatchEvaluatorUnterminatedTimestampFormatPatternTokenEvaluatorInvalidTimestampFormatPatternTokenEvaluatorInvalidTimestampFormatPatternSymbolEvaluatorBindingDoesNotExistMissingHeadersInvalidColumnIndexAdminConfigNotificationTargetsFailedAdminProfilerNotEnabledInvalidDecompressedSizeAddUserInvalidArgumentAddUserValidUTFAdminResourceInvalidArgumentAdminAccountNotEligibleAccountNotEligibleAdminServiceAccountNotFoundPostPolicyConditionInvalidFormatInvalidChecksumLambdaARNInvalidLambdaARNNotFoundInvalidAttributeNameAdminNoAccessKeyAdminNoSecretKeyIAMNotInitializedapiErrCodeEnd"

var _APIErrorCode_index = [...]uint16{0, 4, 16, 25, 39, 53, 67, 81, 94, 112, 129, 144, 161, 174, 186, 208, 228, 254, 268, 289, 306, 321, 344, 361, 379, 396, 420, 435, 456, 474, 486, 506, 523, 546, 567, 579, 597, 618, 646, 676, 697, 720, 746, 783, 813, 846, 871, 903, 933, 962, 987, 1009, 1035, 1057, 1085, 1114, 1148, 1179, 1216, 1240, 1264, 1292, 1318, 1349, 1379, 1388, 1400, 1416, 1429, 1443, 1461, 1481, 1502, 1518, 1529, 1545, 1556, 1584, 1604, 1620, 1648, 1662, 1679, 1699, 1712, 1726, 1739, 1752, 1768, 1785, 1806, 1820, 1841, 1854, 1876, 1899, 1915, 1930, 1945, 1966, 1984, 1999, 2016, 2041, 2059, 2082, 2097, 2116, 2132, 2151, 2172, 2186, 2198, 2211, 2230, 2249, 2259, 2274, 2310, 2341, 2374, 2403, 2415, 2435, 2459, 2483, 2504, 2528, 2547, 2568, 2585, 2595, 2618, 2640, 2666, 2687, 2705, 2732, 2763, 2790, 2811, 2832, 2856, 2881, 2909, 2937, 2953, 2976, 3006, 3017, 3029, 3046, 3061, 3079, 3108, 3125, 3141, 3157, 3175, 3193, 3216, 3237, 3260, 3271, 3287, 3310, 3327, 3355, 3374, 3404, 3424, 3452, 3467, 3485, 3500, 3514, 3549, 3568, 3579, 3592, 3607, 3630, 3656, 3672, 3690, 3708, 3729, 3743, 3760, 3791, 3811, 3832, 3853, 3872, 3891, 3909, 3932, 3956, 3980, 4005, 4040, 4065, 4099, 4132, 4153, 4167, 4186, 4215, 4238, 4265, 4299, 4331, 4361, 4384, 4412, 4444, 4472, 4496, 4520, 4549, 4567, 4584, 4606, 4623, 4641, 4661, 4687, 4703, 4722, 4743, 4747, 4765, 4782, 4808, 4822, 4846, 4867, 4882, 4900, 4923, 4938, 4957, 4974, 4991, 5015, 5042, 5065, 5088, 5105, 5127, 5143, 5163, 5182, 5204, 5225, 5245, 5267, 5291, 5310, 5352, 5373, 5396, 5417, 5448, 5467, 5489, 5509, 5535, 5556, 5578, 5598, 5622, 5645, 5664, 5684, 5706, 5729, 5760, 5798, 5839, 5869, 5883, 5904, 5920, 5942, 5972, 5998, 6026, 6060, 6078, 6101, 6136, 6176, 6218, 6250, 6267, 6292, 6307, 6324, 6334, 6345, 6383, 6437, 6483, 6535, 6583, 6626, 6670, 6698, 6712, 6730, 6766, 6789, 6812, 6834, 6849, 6877, 6900, 6918, 6945, 6977, 6992, 7008, 7025, 7045, 7061, 7077, 7094, 7107}

func (i APIErrorCode) String() string {
	if i < 0 || i >= APIErrorCode(len(_APIErrorCode_index)-1) {
		return "APIErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _APIErrorCode_name[_APIErrorCode_index[i]:_APIErrorCode_index[i+1]]
}