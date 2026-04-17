import { ENV_API_BASE_URL_ERROR } from "@/constants/copy";

const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

if (!apiBaseUrl) {
  throw new Error(ENV_API_BASE_URL_ERROR);
}

export const env = {
  apiBaseUrl
};
