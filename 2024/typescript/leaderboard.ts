import { parseArgs } from "@std/cli/parse-args";
import { join } from "@std/path/join";

export async function fetchLeaderboard(
  id: number,
  year: number = new Date().getFullYear(),
  session?: string
): Promise<any> {
  if (!session) {
    session = Deno.env.get("AOC_SESSION");
  }

  const url = `https://adventofcode.com/${year}/leaderboard/private/view/${id}.json`;
  const response = await fetch(url, {
    headers: {
      Cookie: `session=${session}`,
    },
  });

  return response.json();
}

export async function fileExists(filename: string): Promise<boolean> {
  try {
    const fileInfo = await Deno.stat(filename);
    return fileInfo.isFile;
  } catch (e) {
    if (e instanceof Deno.errors.NotFound) {
      return false;
    }
  }

  return false;
}

export async function downloadLeaderboard(id: number): Promise<any> {
  const leaderboard = await fetchLeaderboard(id);
  await Deno.writeTextFile(
    join(".leaderboard", `${id}.json`),
    JSON.stringify(leaderboard, null, 2)
  );

  return leaderboard;
}

export async function getLeaderboard(id: number): Promise<any> {
  const exists = await fileExists(join(".leaderboard", `${id}.json`));
  if (exists) {
    const leaderboard = JSON.parse(
      await Deno.readTextFile(join(".leaderboard", `${id}.json`))
    );

    // check if leaderboard was fetched in the last 5 minutes
    const now = new Date();
    const lastFetch = new Date(leaderboard.fetched_at);
    if (now.getTime() - lastFetch.getTime() < 5 * 60 * 1000) {
      console.log("using cached dashboard");
      return leaderboard;
    }

    return await downloadLeaderboard(id);
  }

  return await downloadLeaderboard(id);
}

const config = JSON.parse(
  await Deno.readTextFile(join(".leaderboard", "config.json"))
);

const leaderboardNamesToIds: Record<string, number> = {};
for (const [id, info] of Object.entries(config.info)) {
  leaderboardNamesToIds[(info as { [x: string]: string }).name.toLowerCase()] =
    parseInt(id, 10);
}

const args = parseArgs(Deno.args, {
  string: ["ids", "dates", "days", "users", "leaderboards"],
  collect: ["ids", "dates", "days", "users", "leaderboards"],
});

const dashboardIdsFromNames = new Set<number>(
  args.leaderboards.map((name: string) => leaderboardNamesToIds[name.toLowerCase()])
    ?.filter((x?: number) => x)
);

const dashboardIds = new Set<string>(
  config.leaderboards.concat(
    args.ids
      .flatMap((s: string) => s.split(",").map((x: string) => parseInt(x, 10)))
      .concat([...dashboardIdsFromNames])
  )
);

const dates = new Set(
  args.dates.flatMap((s: string) =>
    s.split(",").map((x: string) => parseInt(x, 10))
  )
);

const days = new Set(
  args.days.flatMap((s: string) =>
    s.split(",").map((x: string) => parseInt(x, 10))
  )
);

const users = new Set(args.users.flatMap((s: string) => s.split(",")));

interface LeaderboardLog {
  user: string;
  day: number;
  level: number;
  time: Date;
}

const TEMP_FOLDER = ".leaderboard";

await Deno.mkdir(TEMP_FOLDER, { recursive: true });

for (const id of dashboardIds) {
  const leaderboardName = config.info[id]?.name ?? id.toString();

  console.log(`Leaderboard ${leaderboardName}`);
  const leaderboard = await getLeaderboard(parseInt(id, 10));

  let logEntries: LeaderboardLog[] = [];

  Object.values(leaderboard.members).forEach((member) => {
    const memberName = member.name;
    const memberCompletionDayLevel = member.completion_day_level;

    for (const day of Object.keys(memberCompletionDayLevel)) {
      const dayNumber = parseInt(day, 10);
      const dayLevel = memberCompletionDayLevel[day];

      for (const level of Object.keys(dayLevel)) {
        const levelNumber = parseInt(level, 10);
        const levelInfo = dayLevel[level];

        const time = new Date(levelInfo.get_star_ts * 1000);

        logEntries.push({
          user: memberName,
          day: dayNumber,
          level: levelNumber,
          time,
        });
      }
    }
  });

  logEntries = logEntries.filter((entry) => {
    if (days.size > 0 && !days.has(entry.day)) {
      return false;
    }

    if (dates.size > 0 && !dates.has(entry.time.getDate())) {
      return false;
    }

    if (users.size > 0 && !users.has(entry.user)) {
      return false;
    }

    return true;
  });

  logEntries.sort((a, b) => a.time.getTime() - b.time.getTime());

  logEntries.forEach((entry, idx, entries) => {
    const prevEntry = entries[idx - 1];

    if (idx === 0 || entry.time.getDate() !== prevEntry.time.getDate()) {
      console.log(`  Dec ${entry.time.getDate()}`);
    }

    console.log(
      `    ${entry.time.toLocaleTimeString()} ${entry.user} Day ${entry.day} ${
        entry.level == 1 ? "⭐" : "⭐⭐"
      }`
    );
  });
}
