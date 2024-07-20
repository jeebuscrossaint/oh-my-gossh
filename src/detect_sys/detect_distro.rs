use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

pub fn find_distro() -> io::Result<String> {
    let path = Path::new("/etc/os-release");
    let file = File::open(&path)?;
    let reader = io::BufReader::new(file);

    let distro_name = String::new();

    for line in reader.lines() {
        let line = line?;
        if line.starts_with("ID=") {
            let distro = line.split('=').nth(1).unwrap().trim().to_string();
            println!("Detected distro: {}", distro);
            return Ok(distro); // Return the distro name here
        }
    }

    Ok(distro_name)
}
