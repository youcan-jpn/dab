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
  CurrencyListResResponse,
  NewOrUpdatedCurrencyResResponse,
  NewOrUpdatedCurrencyBody,
  CurrencyId,
} from "../../models";

/**
 * @summary 通貨情報を取得
 */
export const getCurrencies = (
  options?: AxiosRequestConfig
): Promise<AxiosResponse<CurrencyListResResponse>> => {
  return axios.get(`/currencies`, options);
};

export const getGetCurrenciesQueryKey = () => [`/currencies`];

export type GetCurrenciesQueryResult = NonNullable<
  Awaited<ReturnType<typeof getCurrencies>>
>;
export type GetCurrenciesQueryError = AxiosError<unknown>;

export const useGetCurrencies = <
  TData = Awaited<ReturnType<typeof getCurrencies>>,
  TError = AxiosError<unknown>
>(options?: {
  query?: UseQueryOptions<
    Awaited<ReturnType<typeof getCurrencies>>,
    TError,
    TData
  >;
  axios?: AxiosRequestConfig;
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetCurrenciesQueryKey();

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getCurrencies>>> = ({
    signal,
  }) => getCurrencies({ signal, ...axiosOptions });

  const query = useQuery<
    Awaited<ReturnType<typeof getCurrencies>>,
    TError,
    TData
  >({ queryKey, queryFn, ...queryOptions }) as UseQueryResult<TData, TError> & {
    queryKey: QueryKey;
  };

  query.queryKey = queryKey;

  return query;
};

/**
 * @summary 通貨情報を登録
 */
export const postCurrencies = (
  newOrUpdatedCurrencyBody: NewOrUpdatedCurrencyBody,
  options?: AxiosRequestConfig
): Promise<AxiosResponse<NewOrUpdatedCurrencyResResponse>> => {
  return axios.post(`/currencies`, newOrUpdatedCurrencyBody, options);
};

export type PostCurrenciesMutationResult = NonNullable<
  Awaited<ReturnType<typeof postCurrencies>>
>;
export type PostCurrenciesMutationBody = NewOrUpdatedCurrencyBody;
export type PostCurrenciesMutationError = AxiosError<unknown>;

export const usePostCurrencies = <
  TError = AxiosError<unknown>,
  TContext = unknown
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof postCurrencies>>,
    TError,
    { data: NewOrUpdatedCurrencyBody },
    TContext
  >;
  axios?: AxiosRequestConfig;
}) => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof postCurrencies>>,
    { data: NewOrUpdatedCurrencyBody }
  > = (props) => {
    const { data } = props ?? {};

    return postCurrencies(data, axiosOptions);
  };

  return useMutation<
    Awaited<ReturnType<typeof postCurrencies>>,
    TError,
    { data: NewOrUpdatedCurrencyBody },
    TContext
  >(mutationFn, mutationOptions);
};
/**
 * @summary 通貨情報を更新
 */
export const patchCurrenciesCurrencyId = (
  currencyId: CurrencyId,
  newOrUpdatedCurrencyBody: NewOrUpdatedCurrencyBody,
  options?: AxiosRequestConfig
): Promise<AxiosResponse<NewOrUpdatedCurrencyResResponse>> => {
  return axios.patch(
    `/currencies/${currencyId}`,
    newOrUpdatedCurrencyBody,
    options
  );
};

export type PatchCurrenciesCurrencyIdMutationResult = NonNullable<
  Awaited<ReturnType<typeof patchCurrenciesCurrencyId>>
>;
export type PatchCurrenciesCurrencyIdMutationBody = NewOrUpdatedCurrencyBody;
export type PatchCurrenciesCurrencyIdMutationError = AxiosError<unknown>;

export const usePatchCurrenciesCurrencyId = <
  TError = AxiosError<unknown>,
  TContext = unknown
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof patchCurrenciesCurrencyId>>,
    TError,
    { currencyId: CurrencyId; data: NewOrUpdatedCurrencyBody },
    TContext
  >;
  axios?: AxiosRequestConfig;
}) => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof patchCurrenciesCurrencyId>>,
    { currencyId: CurrencyId; data: NewOrUpdatedCurrencyBody }
  > = (props) => {
    const { currencyId, data } = props ?? {};

    return patchCurrenciesCurrencyId(currencyId, data, axiosOptions);
  };

  return useMutation<
    Awaited<ReturnType<typeof patchCurrenciesCurrencyId>>,
    TError,
    { currencyId: CurrencyId; data: NewOrUpdatedCurrencyBody },
    TContext
  >(mutationFn, mutationOptions);
};
