using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.RecoveryServicesBackup.v2023_02_01.ValidateOperation;

internal class Definition : ResourceDefinition
{
    public string Name => "ValidateOperation";
    public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
    {
        new TriggerOperation(),
    };
    public IEnumerable<System.Type> Constants => new List<System.Type>
    {
        typeof(CopyOptionsConstant),
        typeof(OverwriteOptionsConstant),
        typeof(RecoveryModeConstant),
        typeof(RecoveryTypeConstant),
        typeof(RehydrationPriorityConstant),
        typeof(RestoreRequestTypeConstant),
        typeof(SQLDataDirectoryTypeConstant),
        typeof(TargetDiskNetworkAccessOptionConstant),
    };
    public IEnumerable<System.Type> Models => new List<System.Type>
    {
        typeof(AzureFileShareRestoreRequestModel),
        typeof(AzureWorkloadPointInTimeRestoreRequestModel),
        typeof(AzureWorkloadRestoreRequestModel),
        typeof(AzureWorkloadSAPHanaRestoreRequestModel),
        typeof(AzureWorkloadSQLRestoreRequestModel),
        typeof(EncryptionDetailsModel),
        typeof(ExtendedLocationModel),
        typeof(IaasVMRestoreRequestModel),
        typeof(IaasVMRestoreWithRehydrationRequestModel),
        typeof(IdentityBasedRestoreDetailsModel),
        typeof(IdentityInfoModel),
        typeof(RecoveryPointRehydrationInfoModel),
        typeof(RestoreFileSpecsModel),
        typeof(RestoreRequestModel),
        typeof(SQLDataDirectoryMappingModel),
        typeof(SecuredVMDetailsModel),
        typeof(TargetAFSRestoreInfoModel),
        typeof(TargetDiskNetworkAccessSettingsModel),
        typeof(TargetRestoreInfoModel),
        typeof(ValidateIaasVMRestoreOperationRequestModel),
        typeof(ValidateOperationRequestModel),
        typeof(ValidateRestoreOperationRequestModel),
    };
}