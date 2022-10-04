/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package io.cdap.plugin.zuora.objects;

import com.google.gson.annotations.SerializedName;
import io.cdap.cdap.api.data.schema.Schema;
import io.cdap.plugin.zuora.restobjects.annotations.ObjectDefinition;
import io.cdap.plugin.zuora.restobjects.annotations.ObjectFieldDefinition;
import io.cdap.plugin.zuora.restobjects.objects.BaseObject;
import javax.annotation.Nullable;

/** Object name: DebitMemoFromChargeDetailType (DebitMemoFromChargeDetailType). Related objects: */
@SuppressWarnings("unused")
@ObjectDefinition(
    Name = "DebitMemoFromChargeDetailType",
    ObjectType = ObjectDefinition.ObjectDefinitionType.NESTED)
public class DebitMemoFromChargeDetailType extends BaseObject {
  /** Name: amount (amount), Type: number. Options (custom, update, select): false, false, false */
  @Nullable
  @SerializedName("amount")
  @ObjectFieldDefinition(FieldType = Schema.Type.STRING)
  private String amount;

  /**
   * Name: chargeId (chargeId), Type: string. Options (custom, update, select): false, false, false
   */
  @Nullable
  @SerializedName("chargeId")
  @ObjectFieldDefinition(FieldType = Schema.Type.STRING)
  private String chargeId;

  /**
   * Name: comment (comment), Type: string. Options (custom, update, select): false, false, false
   */
  @Nullable
  @SerializedName("comment")
  @ObjectFieldDefinition(FieldType = Schema.Type.STRING)
  private String comment;

  /**
   * Name: description (description), Type: string. Options (custom, update, select): false, false,
   * false
   */
  @Nullable
  @SerializedName("description")
  @ObjectFieldDefinition(FieldType = Schema.Type.STRING)
  private String description;

  /**
   * Name: financeInformation (financeInformation), Type:
   * DebitMemoFromChargeDetailTypeFinanceInformationItem. Options (custom, update, select): false,
   * false, false
   */
  @Nullable
  @SerializedName("financeInformation")
  @ObjectFieldDefinition(FieldType = Schema.Type.STRING)
  private String financeInformation;

  /**
   * Name: memoItemAmount (memoItemAmount), Type: number. Options (custom, update, select): false,
   * false, false
   */
  @Nullable
  @SerializedName("memoItemAmount")
  @ObjectFieldDefinition(FieldType = Schema.Type.STRING)
  private String memoItemAmount;

  /**
   * Name: productRatePlanChargeId (productRatePlanChargeId), Type: string. Options (custom, update,
   * select): false, false, false
   */
  @Nullable
  @SerializedName("productRatePlanChargeId")
  @ObjectFieldDefinition(FieldType = Schema.Type.STRING)
  private String productRatePlanChargeId;

  /**
   * Name: serviceEndDate (serviceEndDate), Type: string. Options (custom, update, select): false,
   * false, false
   */
  @Nullable
  @SerializedName("serviceEndDate")
  @ObjectFieldDefinition(FieldType = Schema.Type.STRING)
  private String serviceEndDate;

  /**
   * Name: serviceStartDate (serviceStartDate), Type: string. Options (custom, update, select):
   * false, false, false
   */
  @Nullable
  @SerializedName("serviceStartDate")
  @ObjectFieldDefinition(FieldType = Schema.Type.STRING)
  private String serviceStartDate;

  @Override
  public void addFields() {
    addCustomField("amount", amount, String.class);
    addCustomField("chargeId", chargeId, String.class);
    addCustomField("comment", comment, String.class);
    addCustomField("description", description, String.class);
    addCustomField("financeInformation", financeInformation, String.class);
    addCustomField("memoItemAmount", memoItemAmount, String.class);
    addCustomField("productRatePlanChargeId", productRatePlanChargeId, String.class);
    addCustomField("serviceEndDate", serviceEndDate, String.class);
    addCustomField("serviceStartDate", serviceStartDate, String.class);
  }
}
