// for that embedded/lightweight linux distro enthusiast that isnt an insane lfs user or the only void linux user in their country

use std::path::Path;

pub fn det_alpine() -> Option<String> {
    let apk_path = "/sbin/apk";
    let apk_static_path = "/sbin/apk.static";

    if Path::new(apk_path).exists() {
        println!("Detected apk at: {}", apk_path);
        Some(apk_path.to_string())
    } else if Path::new(apk_static_path).exists() {
        println!("Detected apk.static at: {}", apk_static_path);
        Some(apk_static_path.to_string())
    } else {
        None
    }
}