import static java.lang.Integer.parseInt;
import static java.lang.System.in;
import static java.lang.System.out;
import static java.util.Comparator.comparing;
import static java.util.Comparator.comparingInt;

void main() {
    var events = new BufferedReader(new InputStreamReader(in)).lines().map(Event::new).sorted(comparing(Event::t)).toList();
    var guards = events.stream().filter(e -> e.id > 0).map(Event::id).distinct().map(id -> new Guard(id, events)).toList();
    var guard1 = guards.stream().max(comparingInt(Guard::asleep)).get();
    var guard2 = guards.stream().max(comparingInt(Guard::mostSleep)).get();
    out.println(guard1.id * guard1.mostMin() + " " + guard2.id * guard2.mostMin());
}

record Event (LocalDateTime t, String action, int id) {
    static final Pattern PTN = Pattern.compile("^\\[(?<ts>\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2})\\] (?<word>\\w+)(?: #(?<id>\\d+))?.*$");
    static final DateTimeFormatter FMT = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm");
    public Event(String s) {
        var m = PTN.matcher(s);
        if (!m.matches()) throw new RuntimeException("not matched: " + s);
        this(LocalDateTime.parse(m.group("ts"), FMT), m.group("word"), m.group("id") != null ? parseInt(m.group("id")) : 0);
    }
}

class Guard {
    int id, asleep, mostMin, mostSleep;

    public Guard(int id, List<Event> events) {
        this.id = id;
        this.asleep = 0;
        var i = 0;
        var minutes = new HashMap<Integer, Integer>();
        while (i < events.size()) {
            if (events.get(i).id == id) {
                i++;
                var falls = 0;
                for (;i < events.size(); i++) {
                    var e = events.get(i);
                    if (e.id > 0) break;
                    switch (e.action()) {
                    case "falls": falls = e.t.getMinute(); break;
                    case "wakes":
                        this.asleep += e.t.getMinute() - falls;
                        for (int m = falls; m < e.t.getMinute(); m++) minutes.merge(m, 1, Integer::sum);
                        break;
                    }
                }
            } else { i++; }
        }
        this.mostMin = this.mostSleep = 0;
        for (var m : minutes.entrySet()) {
            if (m.getValue() > minutes.getOrDefault(this.mostMin, 0)) {
                this.mostMin = m.getKey();
                this.mostSleep = m.getValue();
            }
        }
    }

    int asleep()    { return this.asleep; }
    int mostMin()   { return this.mostMin; }
    int mostSleep() { return this.mostSleep; }
}

