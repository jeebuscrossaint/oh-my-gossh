use std::env;
use std::process::exit;

pub fn find_os() -> &'static str {
    let os = env::consts::OS;

    match os {
        "macos" => "macOS",
        "bsd" => "BSD",
        "windows" => "Windows",
        "linux" => "Linux",
        "android" => "Android",
        _ => {
            println!("Unknown OS, please report what you are running.");
            exit(1);
        },
    }
}