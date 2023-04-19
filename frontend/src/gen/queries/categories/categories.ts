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
  CategoryListResResponse,
  NewOrUpdatedCategoryResResponse,
  NewOrUpdatedCategoryBody,
  CategoryId,
} from "../../models";

/**
 * @summary カテゴリ情報を取得
 */
export const getCategories = (
  options?: AxiosRequestConfig
): Promise<AxiosResponse<CategoryListResResponse>> => {
  return axios.get(`/categories`, options);
};

export const getGetCategoriesQueryKey = () => [`/categories`];

export type GetCategoriesQueryResult = NonNullable<
  Awaited<ReturnType<typeof getCategories>>
>;
export type GetCategoriesQueryError = AxiosError<unknown>;

export const useGetCategories = <
  TData = Awaited<ReturnType<typeof getCategories>>,
  TError = AxiosError<unknown>
>(options?: {
  query?: UseQueryOptions<
    Awaited<ReturnType<typeof getCategories>>,
    TError,
    TData
  >;
  axios?: AxiosRequestConfig;
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetCategoriesQueryKey();

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getCategories>>> = ({
    signal,
  }) => getCategories({ signal, ...axiosOptions });

  const query = useQuery<
    Awaited<ReturnType<typeof getCategories>>,
    TError,
    TData
  >({ queryKey, queryFn, ...queryOptions }) as UseQueryResult<TData, TError> & {
    queryKey: QueryKey;
  };

  query.queryKey = queryKey;

  return query;
};

/**
 * @summary カテゴリ情報を登録
 */
export const postCategories = (
  newOrUpdatedCategoryBody: NewOrUpdatedCategoryBody,
  options?: AxiosRequestConfig
): Promise<AxiosResponse<NewOrUpdatedCategoryResResponse>> => {
  return axios.post(`/categories`, newOrUpdatedCategoryBody, options);
};

export type PostCategoriesMutationResult = NonNullable<
  Awaited<ReturnType<typeof postCategories>>
>;
export type PostCategoriesMutationBody = NewOrUpdatedCategoryBody;
export type PostCategoriesMutationError = AxiosError<unknown>;

export const usePostCategories = <
  TError = AxiosError<unknown>,
  TContext = unknown
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof postCategories>>,
    TError,
    { data: NewOrUpdatedCategoryBody },
    TContext
  >;
  axios?: AxiosRequestConfig;
}) => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof postCategories>>,
    { data: NewOrUpdatedCategoryBody }
  > = (props) => {
    const { data } = props ?? {};

    return postCategories(data, axiosOptions);
  };

  return useMutation<
    Awaited<ReturnType<typeof postCategories>>,
    TError,
    { data: NewOrUpdatedCategoryBody },
    TContext
  >(mutationFn, mutationOptions);
};
/**
 * @summary カテゴリ情報を更新
 */
export const patchCategoriesCategoryId = (
  categoryId: CategoryId,
  newOrUpdatedCategoryBody: NewOrUpdatedCategoryBody,
  options?: AxiosRequestConfig
): Promise<AxiosResponse<NewOrUpdatedCategoryResResponse>> => {
  return axios.patch(
    `/categories/${categoryId}`,
    newOrUpdatedCategoryBody,
    options
  );
};

export type PatchCategoriesCategoryIdMutationResult = NonNullable<
  Awaited<ReturnType<typeof patchCategoriesCategoryId>>
>;
export type PatchCategoriesCategoryIdMutationBody = NewOrUpdatedCategoryBody;
export type PatchCategoriesCategoryIdMutationError = AxiosError<unknown>;

export const usePatchCategoriesCategoryId = <
  TError = AxiosError<unknown>,
  TContext = unknown
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof patchCategoriesCategoryId>>,
    TError,
    { categoryId: CategoryId; data: NewOrUpdatedCategoryBody },
    TContext
  >;
  axios?: AxiosRequestConfig;
}) => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof patchCategoriesCategoryId>>,
    { categoryId: CategoryId; data: NewOrUpdatedCategoryBody }
  > = (props) => {
    const { categoryId, data } = props ?? {};

    return patchCategoriesCategoryId(categoryId, data, axiosOptions);
  };

  return useMutation<
    Awaited<ReturnType<typeof patchCategoriesCategoryId>>,
    TError,
    { categoryId: CategoryId; data: NewOrUpdatedCategoryBody },
    TContext
  >(mutationFn, mutationOptions);
};