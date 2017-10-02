/*
 * Maximum path sum I
 *
 * By starting at the top of the triangle below and moving to adjacent numbers
 * on the row below, the maximum total from top to bottom is 23.
 *
 *    3
 *   7 4
 *  2 4 6
 * 8 5 9 3
 *
 * That is, 3 + 7 + 4 + 9 = 23.
 *
 * Find the maximum total from top to bottom of the triangle below:
 *
 *
 *                            75
 *                          95  64
 *                        17  47  82
 *                      18  35  87  10
 *                    20  04  82  47  65
 *                  19  01  23  75  03  34
 *                88  02  77  73  07  63  67
 *              99  65  04  28  06  16  70  92
 *            41  41  26  56  83  40  80  70  33
 *          41  48  72  33  47  32  37  16  94  29
 *        53  71  44  65  25  43  91  52  97  51  14
 *      70  11  33  28  77  73  17  78  39  68  17  57
 *    91  71  52  38  17  14  91  43  58  50  27  29  48
 *  63  66  04  68  89  53  67  30  73  16  69  87  40  31
 * 04  62  98  27  23  09  70  98  73  93  38  53  60  04  23
 *
 * NOTE: As there are only 16384 routes, it is possible to solve this problem
 * by trying every route. However, Problem 67, is the same challenge with a
 * triangle containing one-hundred rows; it cannot be solved by brute force,
 * and requires a clever method! ;o)
 */

extern crate csv;

#[derive(Debug)]
pub struct Triangle {
    data: Vec<Vec<usize>>
}

impl Triangle {
    pub fn parse(s: &str) -> Triangle {
        let mut data: Vec<Vec<usize>> = vec![];

        // use csv in flexible mode as a quick hack
        let mut rdr = csv::Reader::from_string(s).has_headers(false).delimiter(b' ').flexible(true);
        let mut i = 0;
        let mut last_row_len = 0;
        for row in rdr.decode() {
            let row: Vec<usize> = row.unwrap();
            if i == 0 {
                if row.len() != 1 {
                    panic!("bad initial row");
                }
            } else {
                // i > 0
                if row.len() != last_row_len + 1 {
                    panic!("bad row len: {} vs {}", row.len(), last_row_len);
                }
            }
            last_row_len = row.len();
            i = i + 1;
            data.push(row);
        }

        Triangle { data: data }
    }

    pub fn merge_up_and_reduce(&self) -> usize {
        let mut i = self.data.len() - 1;
        let mut from_row = self.data.get(i).unwrap().clone();
        while i > 0 {
            let mut into_row = self.data.get(i-1).unwrap().clone();
            for i in 0..into_row.len() {
                if from_row[i] >= from_row[i+1] {
                    into_row[i] += from_row[i];
                } else {
                    into_row[i] += from_row[i+1];
                }
            }
            from_row = into_row;
            i = i-1;
        }
        from_row[0]
    }
}

impl PartialEq for Triangle {
    fn eq(&self, other: &Self) -> bool {
        self.data == other.data
    }
}

#[test]
fn test_small() {
    let s = "3\n\
             7 4\n\
             2 4 6\n\
             8 5 9 3\n";

    let t = Triangle::parse(s);
    let want = Triangle { data: vec![vec![3], vec![7, 4], vec![2, 4, 6], vec![8, 5, 9, 3]] };
    assert_eq!(want, t);

    assert_eq!(23, t.merge_up_and_reduce());
}

#[test]
fn test_large() {
    let s = "75\n\
	         95 64\n\
	         17 47 82\n\
	         18 35 87 10\n\
	         20 04 82 47 65\n\
	         19 01 23 75 03 34\n\
             88 02 77 73 07 63 67\n\
	         99 65 04 28 06 16 70 92\n\
	         41 41 26 56 83 40 80 70 33\n\
	         41 48 72 33 47 32 37 16 94 29\n\
	         53 71 44 65 25 43 91 52 97 51 14\n\
	         70 11 33 28 77 73 17 78 39 68 17 57\n\
	         91 71 52 38 17 14 91 43 58 50 27 29 48\n\
	         63 66 04 68 89 53 67 30 73 16 69 87 40 31\n\
	         04 62 98 27 23 09 70 98 73 93 38 53 60 04 23\n";

    let t = Triangle::parse(s);
    assert_eq!(1074, t.merge_up_and_reduce());
}

