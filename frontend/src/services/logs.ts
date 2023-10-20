import axios from "axios";
import dayjs from "dayjs";

const logsApi = axios.create({
  baseURL: 'http://localhost:10658/logs'
})

export interface BotData {
  botQuantity: number;
  playerId: number;
  date: string;
}

export interface BotDataFilters {
  startDate: string;
  endDate: string;
  player?: number;
}

async function fetchLogsData(filters: BotDataFilters): Promise<BotData[]> {
  const formattedFilters = {
    start_date: dayjs(filters.startDate).format("YYYY-MM-DD"),
    end_date: dayjs(filters.endDate).format("YYYY-MM-DD"),
    player_id: filters.player
  }

  const { data } = await logsApi.get('', { params: formattedFilters });
  return data.map((item: any) => ({
    botQuantity: item.bot_quantity,
    playerId: item.player_id,
    date: item.date
  }));
}

export { fetchLogsData };
