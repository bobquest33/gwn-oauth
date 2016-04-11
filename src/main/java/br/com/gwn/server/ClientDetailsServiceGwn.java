package br.com.gwn.server;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.Arrays;

import javax.inject.Inject;
import javax.inject.Named;
import javax.sql.DataSource;

import org.springframework.dao.EmptyResultDataAccessException;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.oauth2.provider.ClientDetails;
import org.springframework.security.oauth2.provider.ClientDetailsService;
import org.springframework.security.oauth2.provider.ClientRegistrationException;
import org.springframework.security.oauth2.provider.client.BaseClientDetails;

@Named
public class ClientDetailsServiceGwn implements ClientDetailsService {

	@Inject
	private DataSource dataSource;

	@Override
	public ClientDetails loadClientByClientId(String clientId) throws ClientRegistrationException {
		JdbcTemplate jdbc = new JdbcTemplate(dataSource);
		
		ClientDetails basic;
		
		try {
			basic = jdbc.queryForObject("select client_id, client_secret, name from apps where client_id = ?", new RowMapper<ClientDetails>() {
				@Override
				public ClientDetails mapRow(ResultSet rs, int idx) throws SQLException {
					BaseClientDetails basic = new BaseClientDetails();
					basic.setClientId(rs.getString("client_id"));
					basic.setClientSecret(rs.getString("client_secret"));
					basic.setScope(Arrays.asList("openid"));
					basic.setAuthorities(Arrays.asList(new SimpleGrantedAuthority("openid")));
					basic.setAuthorizedGrantTypes(Arrays.asList("client_credentials", "password"));
					return basic;
				}
			}, clientId);
		} catch (EmptyResultDataAccessException e) {
			return null;
		}
		
		return basic;
	}

}
