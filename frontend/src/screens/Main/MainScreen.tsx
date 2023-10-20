import { useEffect, useMemo } from "react";
import { useDataLoader } from "../../hooks/useDataFetcher";
import { fetchLogsData } from "../../services/logs";
import dayjs from "dayjs";
import FilterComponent from "./components/MainScreenFilters";
import { useFilters } from "../../hooks/useFilters";
import { Container } from "@mui/material";
import Chart from "./components/Chart";

function MainScreen() {
  const { data, error, isLoading, loadData } = useDataLoader(fetchLogsData);

  const { filters, updateFilters } = useFilters({
    startDate: dayjs().subtract(30, "days").format("YYYY-MM-DD"),
    endDate: dayjs().format("YYYY-MM-DD"),
    player: 0,
  });

  useEffect(() => {
    loadData(filters);
  }, []);

  const unqPlayers = useMemo(() => {
    return Array.from(new Set(data?.map((item) => item.playerId) || []).values());
  }, [data])

  const playersNames = useMemo(() => {
    return Array.from(unqPlayers.values()).map((player) => `Player ${player}`);
  }, [unqPlayers]);

  const aggData = useMemo(() => {
    const dataMap: any = {};
    data?.forEach((item) => {
      dataMap[item.date] = {
        ...dataMap[item.date],
        date: item.date,
        [`Player ${item.playerId}`]: item.botQuantity
      };
    });
    return Object.values(dataMap) as any;
  }, [data]);

  // console.log(data, playersNames, aggData, unqPlayers)
  return (
    <Container>
      <FilterComponent
        filters={filters}
        updateFilters={updateFilters}
        onSubmit={() => loadData(filters)}
        players={unqPlayers}
      />
      <Chart players={playersNames} data={aggData} />
    </Container>
  );
}

export { MainScreen };
