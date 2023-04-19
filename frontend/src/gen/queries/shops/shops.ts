/**
 * Generated by orval v6.12.1 🍺
 * Do not edit manually.
 * dab-api
 * 家計簿アプリ用API
 * OpenAPI spec version: 0.0.1
 */
import axios from "axios";
import type { AxiosRequestConfig, AxiosResponse, AxiosError } from "axios";
import { useQuery, useMutation } from "@tanstack/react-query";
import type {
  UseQueryOptions,
  UseMutationOptions,
  QueryFunction,
  MutationFunction,
  UseQueryResult,
  QueryKey,
} from "@tanstack/react-query";
import type {
  ShopListResResponse,
  NewOrUpdatedShopResResponse,
  NewOrUpdatedShopBody,
  ShopId,
} from "../../models";

/**
 * @summary 店一覧を取得
 */
export const getShops = (
  options?: AxiosRequestConfig
): Promise<AxiosResponse<ShopListResResponse>> => {
  return axios.get(`/shops`, options);
};

export const getGetShopsQueryKey = () => [`/shops`];

export type GetShopsQueryResult = NonNullable<
  Awaited<ReturnType<typeof getShops>>
>;
export type GetShopsQueryError = AxiosError<unknown>;

export const useGetShops = <
  TData = Awaited<ReturnType<typeof getShops>>,
  TError = AxiosError<unknown>
>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof getShops>>, TError, TData>;
  axios?: AxiosRequestConfig;
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetShopsQueryKey();

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getShops>>> = ({
    signal,
  }) => getShops({ signal, ...axiosOptions });

  const query = useQuery<Awaited<ReturnType<typeof getShops>>, TError, TData>({
    queryKey,
    queryFn,
    ...queryOptions,
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryKey;

  return query;
};

/**
 * @summary 新しい店名を登録
 */
export const postShops = (
  newOrUpdatedShopBody: NewOrUpdatedShopBody,
  options?: AxiosRequestConfig
): Promise<AxiosResponse<NewOrUpdatedShopResResponse>> => {
  return axios.post(`/shops`, newOrUpdatedShopBody, options);
};

export type PostShopsMutationResult = NonNullable<
  Awaited<ReturnType<typeof postShops>>
>;
export type PostShopsMutationBody = NewOrUpdatedShopBody;
export type PostShopsMutationError = AxiosError<unknown>;

export const usePostShops = <
  TError = AxiosError<unknown>,
  TContext = unknown
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof postShops>>,
    TError,
    { data: NewOrUpdatedShopBody },
    TContext
  >;
  axios?: AxiosRequestConfig;
}) => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof postShops>>,
    { data: NewOrUpdatedShopBody }
  > = (props) => {
    const { data } = props ?? {};

    return postShops(data, axiosOptions);
  };

  return useMutation<
    Awaited<ReturnType<typeof postShops>>,
    TError,
    { data: NewOrUpdatedShopBody },
    TContext
  >(mutationFn, mutationOptions);
};
/**
 * @summary 店名を更新
 */
export const patchShopsShopId = (
  shopId: ShopId,
  newOrUpdatedShopBody: NewOrUpdatedShopBody,
  options?: AxiosRequestConfig
): Promise<AxiosResponse<NewOrUpdatedShopResResponse>> => {
  return axios.patch(`/shops/${shopId}`, newOrUpdatedShopBody, options);
};

export type PatchShopsShopIdMutationResult = NonNullable<
  Awaited<ReturnType<typeof patchShopsShopId>>
>;
export type PatchShopsShopIdMutationBody = NewOrUpdatedShopBody;
export type PatchShopsShopIdMutationError = AxiosError<unknown>;

export const usePatchShopsShopId = <
  TError = AxiosError<unknown>,
  TContext = unknown
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof patchShopsShopId>>,
    TError,
    { shopId: ShopId; data: NewOrUpdatedShopBody },
    TContext
  >;
  axios?: AxiosRequestConfig;
}) => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof patchShopsShopId>>,
    { shopId: ShopId; data: NewOrUpdatedShopBody }
  > = (props) => {
    const { shopId, data } = props ?? {};

    return patchShopsShopId(shopId, data, axiosOptions);
  };

  return useMutation<
    Awaited<ReturnType<typeof patchShopsShopId>>,
    TError,
    { shopId: ShopId; data: NewOrUpdatedShopBody },
    TContext
  >(mutationFn, mutationOptions);
};