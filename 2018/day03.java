import static java.lang.System.in;
import static java.lang.System.out;

void main() {
    record Pos(int x, int y) {}
    var points = new HashSet<Pos>();
    var overlaps = new HashSet<Pos>();
    var re = Pattern.compile("\\d+");
    var claims = new BufferedReader(new InputStreamReader(in)).lines().map(
        l -> re.matcher(l).results().map(MatchResult::group).mapToInt(Integer::parseInt).toArray()
    ).toList();
    for (var c : claims) {
        for (var x = 0; x < c[3]; x++) {
            for (var y = 0; y < c[4]; y++) {
                var p = new Pos(c[1] + x, c[2] + y);
                if (!points.add(p)) overlaps.add(p);
            }
        }
    }
    out.println(overlaps.size() + " " + claims.stream().filter(
        c -> {
            for (var x = 0; x < c[3]; x++) {
                for (var y = 0; y < c[4]; y++) {
                    if (overlaps.contains(new Pos(c[1] + x, c[2] + y))) return false;
                }
            }
            return true;
        }).mapToInt(c -> c[0]).findFirst().getAsInt());
}
