import { useCallback, useState } from "react";

interface CallbackFunc<ParamsType, ReturnType> {
  (params: ParamsType): Promise<ReturnType>;
}

function useDataLoader<ParamsType, ReturnType>(
  loader: CallbackFunc<ParamsType, ReturnType>
) {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [data, setData] = useState<ReturnType>();

  const loadData = useCallback(async (params: ParamsType) => {
    try {
      setError(null);
      const fetchedData = await loader(params);
      setData(fetchedData);
    } catch(err: any) {
      setError(err.message)
    } finally {
      setIsLoading(false);
    }
  }, [setError, setData, setIsLoading])

  return { isLoading, error, data, loadData };
}

export { useDataLoader }