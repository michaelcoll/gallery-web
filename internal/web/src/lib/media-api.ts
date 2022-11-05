/*
 * Copyright (c) 2022 Michaël COLL.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { getApi } from "@/lib/common-api";
import type { Auth0VueClient } from "@auth0/auth0-vue/src/global";

export interface PhotoApi {
  hash: string;
  dateTime: string;
  iso: number | undefined;
  exposureTime: string | undefined;
  xDimension: number;
  yDimension: number;
  model: string | undefined;
  fNumber: string | undefined;
}

export async function getMediaList(
  auth0Client: Auth0VueClient,
  daemonId: string,
  page: number,
  pageSize: number
): Promise<PhotoApi[]> {
  console.log(daemonId);
  return getApi(auth0Client)
    .then((axiosInstance) =>
      axiosInstance.get<PhotoApi[]>(
        `/api/v1/daemon/${daemonId}/media?page=${page}&pageSize=${pageSize}`
      )
    )
    .then(({ data }) => data);
}
