use std::env;
use std::process::Command;

// window.rs

// TODO: Implement the backend for Windows package managers

pub fn flag_handler() {
    let args: Vec<String> = env::args().collect();
    if args.len() >= 3 {
        let flag = &args[1];
        let pkg_name = &args[2];
        if flag == "i" {
            install_package(pkg_name);
        } else if flag == "r" {
            uninstall_package(pkg_name);
        } else if flag == "u" {
            update_package(pkg_name);
        } else if flag == "se" {
            search_package(pkg_name);
        } 
    }
}

pub fn install_package(package_name: &str) {
    println!("Installing package: {}", package_name);
    let output = Command::new("powershell")
        .arg("-Command")
        .arg(format!("scoop install {}", package_name))
        .output()
        .expect("Failed to execute PowerShell command");
    if output.status.success() {
        println!("Package installed successfully");
    } else {
        let error_message = String::from_utf8_lossy(&output.stderr);
        println!("Failed to install package: {}", error_message);
    }
}

pub fn uninstall_package(package_name: &str) {
    println!("Uninstalling package: {}", package_name);
    let output = Command::new("powershell")
        .arg("-Command")
        .arg(format!("scoop uninstall {}", package_name))
        .output()
        .expect("Failed to execute PowerShell command");
    if output.status.success() {
        println!("Package uninstalled successfully");
    } else {
        let error_message = String::from_utf8_lossy(&output.stderr);
        println!("Failed to uninstall package: {}", error_message);
    }
}

pub fn update_package(package_name: &str) {
    println!("Updating package: {}", package_name);
    let output = Command::new("powershell")
        .arg("-Command")
        .arg(format!("scoop update {}", package_name))
        .output()
        .expect("Failed to execute PowerShell command");
    if output.status.success() {
        println!("Package updated successfully");
    } else {
        let error_message = String::from_utf8_lossy(&output.stderr);
        println!("Failed to update package: {}", error_message);
    }

}

pub fn search_package(package_name: &str) {
    println!("Searching for package: {}", package_name);
    let output = Command::new("powershell")
        .arg("-Command")
        .arg(format!("scoop search {}", package_name))
        .output()
        .expect("Failed to execute PowerShell command");
    if output.status.success() {
        let search_results = String::from_utf8_lossy(&output.stdout);
        println!("{}", search_results);
    } else {
        let error_message = String::from_utf8_lossy(&output.stderr);
        println!("Failed to search for package: {}", error_message);
    }

}

pub fn bucket_mng(bucket_name: &str) {

}

pub fn cache_mng() {

}

pub fn health_check() {

}

pub fn update_all() {

}

pub fn cleanup_cache() {

}

pub fn list_deps() {

}

pub fn create_manifest() {

}

pub fn only_download() {

}

pub fn export_list() {

}

pub fn help() {

}

pub fn pause() {

}

pub fn homepage() {

}

pub fn info() {

}

pub fn import() {

}

pub fn list() {

}

pub fn path_prefix() {

}

pub fn reset() {

}

pub fn shim_manip() {

}

pub fn status_check() {

}

pub fn unpause() {

}

pub fn upgrade() {

}

pub fn virus_check() {

}

pub fn find_app() {
    
}