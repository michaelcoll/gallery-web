import { callExternalApi } from "./external-api.service";

const apiServerUrl = import.meta.env.VITE_API_SERVER_URL;

export const getDaemonList = async (accessToken) => {
  const config = {
    url: `${apiServerUrl}/api/v1/daemon`,
    method: "GET",
    headers: {
      "content-type": "application/json",
      Authorization: `Bearer ${accessToken}`,
    },
  };

  const { data, error } = await callExternalApi({ config });

  return {
    data: data || null,
    error,
  };
};

export const daemonIsAlive = async (daemonId) => {
  const config = {
    url: `${apiServerUrl}/api/v1/daemon/${daemonId}/status`,
    method: "GET",
  };

  const { error } = await callExternalApi({ config });

  return error == null;
};
