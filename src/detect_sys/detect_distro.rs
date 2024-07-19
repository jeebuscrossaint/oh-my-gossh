use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

pub fn find_distro() -> io::Result<()> {
    let path = Path::new("/etc/os-release");
    let file = File::open(&path)?;
    let reader = io::BufReader::new(file);

    for line in reader.lines() {
        let line = line?;
        if line.starts_with("ID=") {
            let distro = line.split('=').nth(1).unwrap();
            println!("Distro: {}", distro);
            break;
        }
    }

    Ok(())
}
