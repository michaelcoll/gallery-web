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

import type dayjs from "dayjs";
import type { PhotoApi } from "@/lib/media-api";
import type { Auth0VueClient } from "@auth0/auth0-vue/src/global";

const apiServerUrl = import.meta.env.VITE_API_SERVER_URL;

export interface GalleryImage {
  largeURL: string;
  thumbnailURL: string;
  width: number;
  height: number;
  date: dayjs.Dayjs;
}

export async function buildImage(
  auth0Client: Auth0VueClient,
  photo: PhotoApi,
  daemonId: string,
  date: dayjs.Dayjs
): Promise<GalleryImage> {
  return mapGalleryImage(photo, daemonId, date);
}

function mapGalleryImage(
  photo: PhotoApi,
  daemonId: string,
  date: dayjs.Dayjs
): GalleryImage {
  return {
    largeURL: `${apiServerUrl}/api/v1/daemon/${daemonId}/media/${photo.hash}/content`,
    thumbnailURL: `${apiServerUrl}/api/v1/daemon/${daemonId}/media/${photo.hash}/thumbnail`,
    width: photo.xDimension,
    height: photo.yDimension,
    date: date,
  };
}
