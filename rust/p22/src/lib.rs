/*
 * Names scores
 *
 * Using names.txt (right click and 'Save Link/Target As...'), a 46K text file
 * containing over five-thousand first names, begin by sorting it into alphabetical
 * order. Then working out the alphabetical value for each name, multiply this value
 * by its alphabetical position in the list to obtain a name score.
 *
 * For example, when the list is sorted into alphabetical order, COLIN, which is worth
 * 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain
 * a score of 938 * 53 = 49714.
 *
 * What is the total of all the name scores in the file?
 */

extern crate csv;

use std::io::Read;

pub fn name_score(name: &str) -> usize {
    let mut score = 0;
    for c in name.chars() {
        let v = (c as u8) - ('A' as u8) + 1;
        score += v as usize;
    }
    score
}

pub fn total_name_scores<T: Read>(r: T) -> usize {
    let mut rdr = csv::Reader::from_reader(r).has_headers(false);
    let mut rows = rdr.decode().collect::<csv::Result<Vec<Vec<String>>>>().unwrap();
    assert_eq!(rows.len(), 1);

    let mut names = rows.swap_remove(0);
    names.sort();

    let mut total_score = 0;
    let mut i = 1;
    for x in &names {
        total_score += i * name_score(x.as_str());
        i += 1;
    }
    total_score
}

#[cfg(test)]
mod test {
    use std::error::Error;
    use std::fs::File;
    use std::path::Path;
    use super::*;

    #[test]
    fn test_name_score() {
        assert_eq!(53, name_score("COLIN"));
    }

    #[test]
    fn test_total_name_scores() {
        let path = Path::new("names.txt");
        let display = path.display();

        let file = match File::open(&path) {
            Err(why) => panic!("couldn't open {}: {}", display, Error::description(&why)),
            Ok(file) => file,
        };

        assert_eq!(871198282, total_name_scores(&file));
    }
}
