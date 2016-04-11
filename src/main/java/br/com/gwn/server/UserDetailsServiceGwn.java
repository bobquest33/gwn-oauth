package br.com.gwn.server;

import java.util.Collections;

import javax.inject.Named;

import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;

@Named
public class UserDetailsServiceGwn implements UserDetailsService {

	@Override
	public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
		User fakeUser = new User("acme", "acme00", Collections.<GrantedAuthority>emptyList());
		return fakeUser;
	}

}
