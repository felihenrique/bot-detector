import { useCallback, useState } from "react";

function useFilters<FiltersType>(initialValues: FiltersType) {
  const [filters, setFilters] = useState<FiltersType>(initialValues);
  const updateFilters = useCallback(
    (item: Partial<FiltersType>) => {
      setFilters((filters) => ({ ...filters, ...item }));
    },
    [setFilters]
  );

  return {
    updateFilters,
    filters,
  };
}

export { useFilters };
