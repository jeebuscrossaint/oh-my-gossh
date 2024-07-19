use std::path::Path;
use std::process::Command;
use std::str;

pub fn det_mac() -> Option<String> {
    let cpu_brand = Command::new("sysctl")
        .args(["-n", "machdep.cpu.brand_string"])
        .output()
        .ok()?
        .stdout;

        let cpu_brand_str = str::from_utf8(&cpu_brand).unwrap_or_default();

        let brew_path = if cpu_brand_str.contains("Apple") {
            "/opt/homebrew/bin/brew"
        } else {
            "/usr/local/bin/brew"
        };

        if Path::new(brew_path).exists() {
            Some(brew_path.to_string())
        } else {
            None
        }
}