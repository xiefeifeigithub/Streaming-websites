package dbops

import "log"

/*
Scheduler服务的业务逻辑
1. user -> api service ->delete video
2. api server -> scheduler -> write video deletion record
3. timer
4. timer -> runner -> read wvdr -> exec -> delete video from folder
 */

func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) VALUES (?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error: %v", err)
		return err
	}

	defer stmtIns.Close()
	return nil
}