package Database

import (
	"backend/Class/Utils"
	"backend/Entity"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

// <editor-fold desc="For Table: charts">

func GetAllCharts() (*sql.Rows, error) {
	return DB.Query(`SELECT * FROM charts`)
}

func GetALlChartsOrderedByName() (*sql.Rows, error) {
	return DB.Query(`SELECT * FROM charts GROUP BY name;`)
}

/*
GetChartByFilename Retrieve a Chart by his filename in path
*/
func GetChartByFilename(filename string) *sql.Row {
	return DB.QueryRow(fmt.Sprintf("SELECT * FROM charts WHERE path = '%s'", filename))
}

/*
GetChartByCriteria Get a unique chart by is name and version and path (location in chart directory (sub-folder/project folder))
Rule : There may be two charts with the same name and version but different because they are located in another folder.
*/
func GetChartByCriteria(chart Entity.ChartDTO) (*sql.Rows, error) {
	// Maybe calculate the digest by sha256 the Chart.yaml of a chart and compare with chart's digest in db

	var id int
	err := DB.QueryRow(`
        SELECT id 
        FROM charts
        WHERE name = ? 
        AND version = ? 
        AND path = ?
        LIMIT 1;
    `, "tes-nginx", "1.0.0", Utils.NullToString(chart.Path)).Scan(&id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Aucun chart trouvé avec ces critères.")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Chart ID: %d\n", id)
	}

	return DB.Query("SELECT id from charts")
}

// </editor-fold>

// <editor-fold desc="For Table: registry">

func GetInfo() *sql.Row {
	return DB.QueryRow(`SELECT * FROM registry;`)
}

// </editor-fold>
