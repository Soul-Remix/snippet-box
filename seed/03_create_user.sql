CREATE USER 'snippet' @'%';
GRANT SELECT,
    INSERT,
    UPDATE,
    DELETE ON snippetbox.* TO 'snippet' @'%';
ALTER USER 'snippet' @'%' IDENTIFIED BY 'pass';