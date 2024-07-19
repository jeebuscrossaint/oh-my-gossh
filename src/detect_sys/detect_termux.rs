use std::env;

pub fn det_termux() -> bool {
    match env::var("PREFIX") {
        Ok(val) => !val.is_empty(),
        Err(_) => false,
    }
}