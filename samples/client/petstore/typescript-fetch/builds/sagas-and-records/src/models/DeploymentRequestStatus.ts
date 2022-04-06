/* tslint:disable */
/* eslint-disable */
/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * Status of the deployment request
 * @export
 */
export const DeploymentRequestStatus = {
    New: 'New',
    Prepared: 'Prepared',
    Printed: 'Printed',
    Tested: 'Tested',
    Completed: 'Completed',
    Cancelled: 'Cancelled',
    Promoted: 'Promoted',
    Assigned: 'Assigned',
    Ready: 'Ready',
    Packaged: 'Packaged',
    Pairing: 'Pairing',
    Paired: 'Paired'
} as const;

export type DeploymentRequestStatus = typeof DeploymentRequestStatus[keyof typeof DeploymentRequestStatus];


export function DeploymentRequestStatusFromJSON(json: any): DeploymentRequestStatus {
    return DeploymentRequestStatusFromJSONTyped(json, false);
}

export function DeploymentRequestStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploymentRequestStatus {
    return json as DeploymentRequestStatus;
}

export function DeploymentRequestStatusToJSON(value?: DeploymentRequestStatus | null): any {
    return value as any;
}

