drop schema cac_master;
CREATE SCHEMA cac_master;

CREATE TABLE IF NOT EXISTS cac_master.kubernetes_provision(
	kubernetes_provision_id INTEGER GENERATED ALWAYS AS IDENTITY,
	cluster_name CHARACTER VARYING(64),
	node_count INTEGER,
	PRIMARY KEY (kubernetes_provision_id)
)

insert into cac_master.kubernetes_provision 
values (default, 'fsdd2f-provision', 4);

insert into cac_master.kubernetes_provision 
values (default, '12h343-provision', 6);

insert into cac_master.kubernetes_provision 
values (default, 'fdsh31-provision', 1);

insert into cac_master.kubernetes_provision 
values (default, 'lihn23-provision', 3);

-- CRUD Operations 
-- getKubenetesProvision()
CREATE OR REPLACE FUNCTION cac_master.list_clusters()
RETURNS TABLE (kubernetes_provision_id INTEGER, cluster_name CHARACTER VARYING, node_count INTEGER) AS
$BODY$
BEGIN
    return query
        SELECT * FROM cac_master.kubernetes_provision;
END;
$BODY$ LANGUAGE plpgsql;

SELECT cac_master.list_clusters()


-- getClusterDetail()
CREATE OR REPLACE FUNCTION cac_master.get_cluster_detail(prov_id INTEGER)
RETURNS TABLE (provId INTEGER, clusterName CHARACTER VARYING(64), nodeCount INTEGER) AS
$BODY$
BEGIN
	RETURN QUERY
		SELECT
			kubernetes_provision_id,
			cluster_name,
			node_count
		FROM cac_master.kubernetes_provision
		WHERE kubernetes_provision_id=prov_id;
END;
$BODY$ LANGUAGE plpgsql;

-- drop function cac_master.get_cluster_detail;
-- drop function cac_master_delete_cluster;


select cac_master.get_cluster_detail(1);


CREATE OR REPLACE FUNCTION cac_master.update_cluster(prov_id INTEGER,clust_name CHARACTER VARYING(64), nod_count INTEGER)
RETURNS VOID AS
$$
BEGIN
	UPDATE cac_master.kubernetes_provision
	SET cluster_name=clust_name, node_count=nod_count
	WHERE kubernetes_provision_id=prov_id;
END;
$$ LANGUAGE plpgsql;

DROP FUNCTION cac_master.update_cluster(integer,character varying,integer);


select cac_master.update_cluster(2,'UPDDF32-provision', 2);

CREATE OR REPLACE FUNCTION cac_master.delete_cluster(prov_id INTEGER) 
RETURNS VOID AS
$$
BEGIN
	DELETE
	FROM cac_master.kubernetes_provision
	WHERE kubernetes_provision_id=prov_id;
END;
$$ LANGUAGE plpgsql;

select cac_master.delete_cluster(1);
